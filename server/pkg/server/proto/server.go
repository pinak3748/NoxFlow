package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type LogStreamingServer struct {
	UnimplementedLogStreamingServiceServer
}

type UsageStreamingServer struct {
	UnimplementedUsageStreamingServiceServer
}

// StreamLogs implements the bidirectional streaming RPC for logs
func (s *LogStreamingServer) StreamLogs(stream LogStreamingService_StreamLogsServer) error {
	for {
		// Receive log data from client
		logData, err := stream.Recv()
		if err != nil {
			return fmt.Errorf("error receiving log data: %v", err)
		}

		// Process the received log data
		log.Printf("Received log from container %s: %s",
			logData.Metadata.ContainerName,
			logData.Log)

		// Send response back to client
		if err := stream.Send(&LogResponse{
			Message: fmt.Sprintf("Received log from container %s", logData.Metadata.ContainerName),
		}); err != nil {
			return fmt.Errorf("error sending response: %v", err)
		}
	}
}

// StreamUsage implements the bidirectional streaming RPC for container usage stats
func (s *UsageStreamingServer) StreamUsage(stream UsageStreamingService_StreamUsageServer) error {
	for {
		// Receive usage stats from client
		usageStats, err := stream.Recv()
		if err != nil {
			return fmt.Errorf("error receiving usage stats: %v", err)
		}

		// Process the received usage stats
		log.Printf("Received usage stats from container %s: CPU: %.2f%%, Memory: %.2f%%",
			usageStats.ContainerId,
			usageStats.CpuPercent,
			usageStats.MemoryPercent)

		// Send response back to client
		if err := stream.Send(&UsageResponse{
			Message: fmt.Sprintf("Received usage stats from container %s", usageStats.ContainerId),
		}); err != nil {
			return fmt.Errorf("error sending response: %v", err)
		}
	}
}

// StartServer initializes and starts the gRPC server
func StartServer(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register our services
	RegisterLogStreamingServiceServer(s, &LogStreamingServer{})
	RegisterUsageStreamingServiceServer(s, &UsageStreamingServer{})

	log.Printf("Starting gRPC server on port %d", port)
	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}
