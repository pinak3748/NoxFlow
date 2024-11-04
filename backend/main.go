package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type ContainerLogMetadata struct {
	ContainerID   string
	ContainerName string
	Image         string
	State         string
	LogPath       string
	LogDriver     string
}

type LogData struct {
	Metadata ContainerLogMetadata
	Log      string
}

var (
	clients   = make(map[chan LogData]bool)
	clientsMu sync.RWMutex
)

func handleLogStream(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Handle incoming logs
		var logData LogData
		if err := json.NewDecoder(r.Body).Decode(&logData); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Broadcast to all connected clients
		clientsMu.RLock()
		for client := range clients {
			client <- logData
		}
		clientsMu.RUnlock()

		w.WriteHeader(http.StatusOK)
		return
	}

	// Handle SSE connection
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Create a channel for this client
	logChan := make(chan LogData)

	clientsMu.Lock()
	clients[logChan] = true
	clientsMu.Unlock()

	// Clean up when the client disconnects
	defer func() {
		clientsMu.Lock()
		delete(clients, logChan)
		close(logChan)
		clientsMu.Unlock()
	}()

	// Stream logs to the client
	for logData := range logChan {
		data, err := json.Marshal(logData)
		if err != nil {
			continue
		}
		fmt.Fprintf(w, "data: %s\n\n", data)
		w.(http.Flusher).Flush()
	}
}

func main() {
	http.HandleFunc("/logs/stream", handleLogStream)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
