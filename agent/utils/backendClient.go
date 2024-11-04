package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type BackendClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

type LogData struct {
	Metadata ContainerLogMetadata
	Log      string
}

type ContainerLogMetadata struct {
	ContainerID   string
	ContainerName string
	Image         string
	State         string
	LogPath       string
	LogDriver     string
}

func NewBackendClient(baseURL string) *BackendClient {
	return &BackendClient{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (c *BackendClient) SendLog(metadata *ContainerLogMetadata, logLine string) error {
	maxRetries := 3
	backoff := time.Second

	logData := LogData{
		Metadata: *metadata,
		Log:      logLine,
	}

	jsonData, err := json.Marshal(logData)
	if err != nil {
		return fmt.Errorf("error marshaling log data: %v", err)
	}

	for attempt := 1; attempt <= maxRetries; attempt++ {
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/logs/stream", c.BaseURL), bytes.NewBuffer(jsonData))
		if err != nil {
			return fmt.Errorf("error creating request: %v", err)
		}

		req.Header.Set("Content-Type", "application/json")

		resp, err := c.HTTPClient.Do(req)
		if err != nil {
			if attempt == maxRetries {
				return fmt.Errorf("error sending request after %d attempts: %v", maxRetries, err)
			}
			log.Printf("Attempt %d failed, retrying in %v...", attempt, backoff)
			time.Sleep(backoff)
			backoff *= 2 // Exponential backoff
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			if attempt == maxRetries {
				return fmt.Errorf("unexpected status code after %d attempts: %d", maxRetries, resp.StatusCode)
			}
			log.Printf("Attempt %d failed with status %d, retrying in %v...", attempt, resp.StatusCode, backoff)
			time.Sleep(backoff)
			backoff *= 2
			continue
		}

		return nil // Success
	}

	return fmt.Errorf("failed to send log after %d attempts", maxRetries)
}
