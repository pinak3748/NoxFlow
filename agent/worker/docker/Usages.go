package docker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/noxflow-agent/utils"
)

type ContainerUsageStats struct {
	ContainerID string    `json:"container_id"`
	Timestamp   time.Time `json:"timestamp"`
	// Add the usage data is in bytes
	CPUPercent     float64 `json:"cpu_percent"`
	CPUUsage       uint64  `json:"cpu_usage"`
	SystemCPUUsage uint64  `json:"system_cpu_usage"`
	MemoryUsage    uint64  `json:"memory_usage"`
	MemoryLimit    uint64  `json:"memory_limit"`
	MemoryPercent  float64 `json:"memory_percent"`
	MemoryCache    uint64  `json:"memory_cache"`
}

func GetDockerContainerUsage(containerID string, wg *sync.WaitGroup, stream bool) {
	defer wg.Done()

	log.Printf("Starting container stats collection for: %s", containerID)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if !stream {
		// Get one-time stats
		stats, err := getContainerStats(ctx, containerID)
		if err != nil {
			log.Printf("Error getting one-time stats: %v", err)
			return
		}
		processStats(stats)
		return
	}

	// Create new context for streaming
	streamCtx, streamCancel := context.WithCancel(context.Background())
	defer streamCancel()

	// Get streaming stats
	containerStats, err := utils.ContainerStats(streamCtx, containerID, true)
	if err != nil {
		log.Printf("Error initializing stats stream for container %s: %v", containerID, err)
		return
	}
	defer containerStats.Body.Close()

	decoder := json.NewDecoder(containerStats.Body)

	// Start streaming loop
	for {
		select {
		case <-streamCtx.Done():
			log.Printf("Stopping stats stream for container %s: context cancelled", containerID)
			return
		default:
			var statsJSON container.StatsResponse
			if err := decoder.Decode(&statsJSON); err != nil {
				log.Printf("Error decoding stats for container %s: %v", containerID, err)
				return
			}

			stats := extractStats(containerID, &statsJSON)
			processStats(stats)
		}
	}
}

func getContainerStats(ctx context.Context, containerID string) (*ContainerUsageStats, error) {
	containerStats, err := utils.ContainerStats(ctx, containerID, false)
	if err != nil {
		return nil, fmt.Errorf("failed to get container stats: %v", err)
	}
	defer containerStats.Body.Close()

	var statsJSON container.StatsResponse
	if err := json.NewDecoder(containerStats.Body).Decode(&statsJSON); err != nil {
		return nil, fmt.Errorf("failed to decode stats JSON: %v", err)
	}

	return extractStats(containerID, &statsJSON), nil
}

func extractStats(containerID string, statsJSON *container.StatsResponse) *ContainerUsageStats {
	stats := &ContainerUsageStats{
		ContainerID:    containerID,
		Timestamp:      time.Now(),
		CPUUsage:       statsJSON.CPUStats.CPUUsage.TotalUsage,
		SystemCPUUsage: statsJSON.CPUStats.SystemUsage,
		MemoryUsage:    statsJSON.MemoryStats.Usage,
		MemoryLimit:    statsJSON.MemoryStats.Limit,
		MemoryCache:    statsJSON.MemoryStats.Stats["cache"],
	}

	// Calculate CPU percent
	if len(statsJSON.CPUStats.CPUUsage.PercpuUsage) > 0 {
		cpuDelta := float64(statsJSON.CPUStats.CPUUsage.TotalUsage - statsJSON.PreCPUStats.CPUUsage.TotalUsage)
		systemDelta := float64(statsJSON.CPUStats.SystemUsage - statsJSON.PreCPUStats.SystemUsage)

		if systemDelta > 0.0 && cpuDelta > 0.0 {
			stats.CPUPercent = (cpuDelta / systemDelta) * float64(len(statsJSON.CPUStats.CPUUsage.PercpuUsage)) * 100.0
		}
	}

	// Calculate Memory percent
	if stats.MemoryLimit != 0 {
		stats.MemoryPercent = (float64(stats.MemoryUsage) / float64(stats.MemoryLimit)) * 100.0
	}

	return stats
}

// processStats handles the stats data
func processStats(stats *ContainerUsageStats) {
	log.Printf("Container %s - CPU: %.2f%%, Memory: %.2f%%",
		stats.ContainerID, stats.CPUPercent, stats.MemoryPercent)

	// TODO: Implement server sending logic here
}
