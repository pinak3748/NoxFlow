package server

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/nox/noxflow/server-gRPC/utils"
	"google.golang.org/grpc"
)

type LogStreamingServer struct {
	UnimplementedLogStreamingServiceServer
	dbClient *utils.DatabaseClient
}

type UsageStreamingServer struct {
	UnimplementedUsageStreamingServiceServer
	dbClient *utils.DatabaseClient
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

		cleanedLog := strings.Replace(logData.Log, "\x00", "", -1)

		// Save to database
		err = s.dbClient.AddLog(&utils.LogData{
			Timestamp:     time.Now(),
			ContainerName: logData.Metadata.ContainerName,
			LogMessage:    cleanedLog,
		})
		if err != nil {
			log.Printf("Error saving log to database: %v", err)
		}

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

		// Save to database
		err = s.dbClient.AddUsage(&utils.UsageData{
			Timestamp:     time.Now(),
			ContainerID:   usageStats.ContainerId,
			CPUPercent:    usageStats.CpuPercent,
			MemoryPercent: usageStats.MemoryPercent,
		})
		if err != nil {
			log.Printf("Error saving usage stats to database: %v", err)
		}

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

	dbClient, err := utils.NewDatabaseClient(
		"postgres://postgres:password@localhost:55000/postgres?sslmode=disable",
		1000,          // batch size
		5*time.Second, // flush interval
	)
	if err != nil {
		log.Fatalf("Failed to initialize database client: %v", err)
	}
	defer dbClient.Close()

	// Register our services with the database client
	RegisterLogStreamingServiceServer(s, &LogStreamingServer{
		dbClient: dbClient,
	})
	RegisterUsageStreamingServiceServer(s, &UsageStreamingServer{
		dbClient: dbClient,
	})
	log.Printf("Starting gRPC server on port %d", port)
	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}
