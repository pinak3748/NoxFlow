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

	dockerContainers, err := utils.DockerListContainers(ctx)
	if err != nil {
		log.Fatalf("Error listing containers: %v", err)
	}

	var wg sync.WaitGroup

	if len(dockerContainers) == 0 {
		log.Println("No containers found!!!!")
	} else {
		for _, container := range dockerContainers {
			wg.Add(1)
			log.Printf("Container ID: %s", container.ID)
			go docker.ContainerLogs(container.ID, &wg)
		}
	}

	wg.Wait()

}
