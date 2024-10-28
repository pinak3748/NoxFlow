package docker

import (
	"bufio"
	"context"
	"log"
	"sync"

	"github.com/docker/docker/api/types/container"
	pb "github.com/noxflow-agent/pkg/proto"
	"github.com/noxflow-agent/utils"
)

// GetDockerContainerLogs streams logs from a Docker container and sends them to the server
func GetDockerContainerLogs(containerID string, monitorClient *utils.MonitorClient, wg *sync.WaitGroup) {
	defer wg.Done()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	containerInfo, err := utils.DockerClient.ContainerInspect(ctx, containerID)
	if err != nil {
		log.Printf("Error inspecting container %s: %v", containerID, err)
		return
	}

	// Construct the metadata for the logs
	metadata := &pb.ContainerLogMetadata{
		ContainerId:   containerID,
		ContainerName: containerInfo.Name,
		Image:         containerInfo.Config.Image,
		State:         containerInfo.State.Status,
		LogPath:       containerInfo.LogPath,
		LogDriver:     containerInfo.HostConfig.LogConfig.Type,
	}

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

	// Use a scanner to read logs line by line
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		logLine := scanner.Text()

		// Send the log line to the server
		_, err := monitorClient.SendLog(metadata, logLine)
		if err != nil {
			log.Printf("Error sending log for container %s: %v", containerID, err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading logs for container %s: %v", containerID, err)
	}
}
