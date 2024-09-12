package main

import (
	"log"

	"github.com/chub-es/go-link-shortener/config"
	"github.com/chub-es/go-link-shortener/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config load error: %s", err)
	}

	app.Run(cfg)
}
