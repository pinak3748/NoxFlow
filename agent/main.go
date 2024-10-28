package main

import (
	"context"
	"log"
	"sync"

	"github.com/noxflow-agent/utils"
	"github.com/noxflow-agent/worker/docker"
)

func main() {

	// Initialize Docker client
	err := utils.InitDockerClient()
	if err != nil {
		log.Fatalf("Failed to initialize Docker client: %v", err)
	}

	// Create a context
	ctx := context.Background()

	// List Docker containers
	dockerContainers, err := utils.DockerListContainers(ctx)
	if err != nil {
		log.Fatalf("Error listing containers: %v", err)
	}

	// Initialize monitorClient
	monitorClient, err := utils.NewMonitorClient("localhost:8888", 5)
	if err != nil {
		log.Fatalf("Failed to create monitor client: %v", err)
	}
	defer monitorClient.Close()

	var wg sync.WaitGroup

	if len(dockerContainers) == 0 {
		log.Println("No containers found!!!!")
	} else {
		for _, container := range dockerContainers {
			wg.Add(1)
			log.Printf("Container ID: %s", container.ID)
			// Start streaming logs
			go docker.GetDockerContainerLogs(container.ID, monitorClient, &wg)
			// If you also want to start streaming usage, uncomment the next line
			// go docker.GetDockerContainerUsage(container.ID, &wg, false)
		}
	}

	wg.Wait()

}
