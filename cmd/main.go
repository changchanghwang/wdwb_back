package main

import (
	"log"

	"github.com/changchanghwang/wdwb_back/cmd/di"
	"github.com/changchanghwang/wdwb_back/internal/config"
)

func main() {
	server, err := di.InitializeServer()
	if err != nil {
		log.Fatalf("failed to initialize server: %v", err)
	}

	if err := server.Run(":" + config.Port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
