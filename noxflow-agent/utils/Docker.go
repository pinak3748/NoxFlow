package utils

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

var DockerVersion string
var DockerClient *client.Client

func InitDockerClient() error {
	var err error
	DockerClient, err = client.NewClientWithOpts(client.WithVersion("1.47"), client.WithAPIVersionNegotiation())
	return err
}

func DockerListContainers(ctx context.Context) ([]types.Container, error) {
	// Get all the containers (running and stopped)
	options := container.ListOptions{All: true}
	return DockerClient.ContainerList(ctx, options)
}

func DockerInspect(ctx context.Context, containerID string) (types.ContainerJSON, error) {
	return DockerClient.ContainerInspect(ctx, containerID)

}
