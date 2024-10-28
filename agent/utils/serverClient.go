package utils

import (
	"context"
	"fmt"
	"sync"
	"time"

	pb "github.com/noxflow-agent/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// MonitorClient handles the gRPC connections and streams for monitoring
type MonitorClient struct {
	connections       []*grpc.ClientConn
	logClients        []pb.LogStreamingServiceClient
	usageClients      []pb.UsageStreamingServiceClient
	logStreams        []pb.LogStreamingService_StreamLogsClient
	usageStreams      []pb.UsageStreamingService_StreamUsageClient
	ctx               context.Context
	cancel            context.CancelFunc
	mu                sync.Mutex
	reconnectInterval time.Duration
	currentConnIndex  int
}

// NewMonitorClient creates a new instance of MonitorClient with multiple connections
func NewMonitorClient(serverAddr string, numConnections int) (*MonitorClient, error) {
	ctx, cancel := context.WithCancel(context.Background())

	client := &MonitorClient{
		connections:       make([]*grpc.ClientConn, numConnections),
		logClients:        make([]pb.LogStreamingServiceClient, numConnections),
		usageClients:      make([]pb.UsageStreamingServiceClient, numConnections),
		logStreams:        make([]pb.LogStreamingService_StreamLogsClient, numConnections),
		usageStreams:      make([]pb.UsageStreamingService_StreamUsageClient, numConnections),
		ctx:               ctx,
		cancel:            cancel,
		reconnectInterval: 5 * time.Second,
	}

	// Initialize all connections
	for i := 0; i < numConnections; i++ {
		conn, err := grpc.DialContext(ctx, serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			client.Close()
			return nil, fmt.Errorf("failed to connect to server: %v", err)
		}
		client.connections[i] = conn
		client.logClients[i] = pb.NewLogStreamingServiceClient(conn)
		client.usageClients[i] = pb.NewUsageStreamingServiceClient(conn)
	}

	return client, nil
}

// getNextConnection returns the next connection index in a round-robin fashion
func (c *MonitorClient) getNextConnection() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.currentConnIndex = (c.currentConnIndex + 1) % len(c.connections)
	return c.currentConnIndex
}

// initLogStream initializes the log streaming connection for a specific index
func (c *MonitorClient) initLogStream(index int) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.logStreams[index] != nil {
		return nil
	}

	stream, err := c.logClients[index].StreamLogs(c.ctx)
	if err != nil {
		return fmt.Errorf("failed to initialize log stream: %v", err)
	}

	c.logStreams[index] = stream
	return nil
}

// initUsageStream initializes the usage streaming connection for a specific index
func (c *MonitorClient) initUsageStream(index int) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.usageStreams[index] != nil {
		return nil
	}

	stream, err := c.usageClients[index].StreamUsage(c.ctx)
	if err != nil {
		return fmt.Errorf("failed to initialize usage stream: %v", err)
	}

	c.usageStreams[index] = stream
	return nil
}

// SendLog sends container log data to the server using round-robin connection selection
func (c *MonitorClient) SendLog(metadata *pb.ContainerLogMetadata, logData string) (*pb.LogResponse, error) {
	connIndex := c.getNextConnection()
	if err := c.initLogStream(connIndex); err != nil {
		return nil, err
	}

	log := &pb.LogData{
		Metadata: metadata,
		Log:      logData,
	}

	c.mu.Lock()
	stream := c.logStreams[connIndex]
	c.mu.Unlock()

	err := stream.Send(log)
	if err != nil {
		c.mu.Lock()
		c.logStreams[connIndex] = nil
		c.mu.Unlock()
		return nil, fmt.Errorf("failed to send log data: %v", err)
	}

	response, err := stream.Recv()
	if err != nil {
		c.mu.Lock()
		c.logStreams[connIndex] = nil
		c.mu.Unlock()
		return nil, fmt.Errorf("failed to receive log response: %v", err)
	}

	return response, nil
}

// SendUsageStats sends container usage statistics using round-robin connection selection
func (c *MonitorClient) SendUsageStats(stats *pb.ContainerUsageStats) (*pb.UsageResponse, error) {
	connIndex := c.getNextConnection()
	if err := c.initUsageStream(connIndex); err != nil {
		return nil, err
	}

	c.mu.Lock()
	stream := c.usageStreams[connIndex]
	c.mu.Unlock()

	err := stream.Send(stats)
	if err != nil {
		c.mu.Lock()
		c.usageStreams[connIndex] = nil
		c.mu.Unlock()
		return nil, fmt.Errorf("failed to send usage stats: %v", err)
	}

	response, err := stream.Recv()
	if err != nil {
		c.mu.Lock()
		c.usageStreams[connIndex] = nil
		c.mu.Unlock()
		return nil, fmt.Errorf("failed to receive usage response: %v", err)
	}

	return response, nil
}

// Close closes all connections and streams
func (c *MonitorClient) Close() error {
	c.cancel()
	for i := range c.connections {
		if c.logStreams[i] != nil {
			c.logStreams[i].CloseSend()
		}
		if c.usageStreams[i] != nil {
			c.usageStreams[i].CloseSend()
		}
		if c.connections[i] != nil {
			c.connections[i].Close()
		}
	}
	return nil
}
