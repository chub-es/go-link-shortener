package main

import (
	"log"
	"os"

	"github.com/chub-es/go-link-shortener/config"
	"github.com/chub-es/go-link-shortener/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config load error: %s", err)
	}

	if cfg.App.TimeZone != "" {
		os.Setenv("TZ", cfg.App.TimeZone)
	}

	app.Run(cfg)
}
