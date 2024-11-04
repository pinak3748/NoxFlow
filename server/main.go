package main

import (
	"log"

	server "github.com/nox/noxflow/server-gRPC/pkg/server/proto"
)

func main() {

	// Start the server
	if err := server.StartServer(8888); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
