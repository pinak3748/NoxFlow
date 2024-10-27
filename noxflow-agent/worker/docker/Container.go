package docker

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"sync"

	"github.com/docker/docker/api/types/container"
	"github.com/noxflow-agent/utils"
)

type ContainerLogMessage struct {
	Metadata struct {
		ContainerID   string `json:"container_id"`
		ContainerName string `json:"container_name"`
		Image         string `json:"image"`
		State         string `json:"state"`
		LogPath       string `json:"log_path"`
		LogDriver     string `json:"log_driver"`
	} `json:"metadata"`
	Log string `json:"log"`
}

func ContainerLogs(containerID string, wg *sync.WaitGroup) {
	defer wg.Done()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	containerInfo, err := utils.DockerClient.ContainerInspect(ctx, containerID)
	if err != nil {
		log.Printf("Error inspecting container %s: %v", containerID, err)
		return
	}

	logMessage := ContainerLogMessage{}
	logMessage.Metadata.ContainerID = containerID
	logMessage.Metadata.ContainerName = containerInfo.Name
	logMessage.Metadata.Image = containerInfo.Config.Image
	logMessage.Metadata.State = containerInfo.State.Status
	logMessage.Metadata.LogPath = containerInfo.LogPath
	logMessage.Metadata.LogDriver = containerInfo.HostConfig.LogConfig.Type

	reader, err := utils.DockerClient.ContainerLogs(ctx, containerID, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
		Timestamps: true,
		Details:    true,
		Tail:       "0",
	})

	if err != nil {
		log.Printf("Error retrieving logs for container %s: %v", containerID, err)
		return
	}
	defer reader.Close()

	// Loop over the logs
	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Error reading logs for container %s: %v", containerID, err)
			return
		}

		logMessage.Log = string(buf[:n])

		jsonData, err := json.Marshal(logMessage)
		if err != nil {
			log.Printf("Error marshaling log message for container %s: %v", containerID, err)
			continue
		}

		// TODO: Send jsonData to the server
		fmt.Printf("%s\n", string(jsonData))
	}

}
