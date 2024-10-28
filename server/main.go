package main

import (
	"log"

	server "github.com/noxflow-server/pkg/server/proto"
)

func main() {
	if err := server.StartServer(8888); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
