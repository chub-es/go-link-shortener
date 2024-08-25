package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		HTTP
		Log
	}

	// HTTP -.
	HTTP struct {
		Port string `env:"HTTP_PORT" env-default:"8080"`
	}

	// Log -.
	Log struct {
		Level string `env:"LOG_LEVEL" env-default:"debug"`
	}
)

func New() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}
	log.Fatalf(cfg.Port)
	return cfg, nil
}
