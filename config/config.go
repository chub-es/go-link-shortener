package config

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type (
	// Config -.
	Config struct {
		HTTP `mapstructure:",squash"`
		PG   `mapstructure:",squash"`
	}

	// HTTP -.
	HTTP struct {
		Mode           string        `mapstructure:"GIN_MODE"`
		Port           string        `mapstructure:"HTTP_PORT" validate:"required"`
		ReadTimeout    time.Duration `mapstructure:"HTTP_READ_TIMEOUT"`
		WriteTimeout   time.Duration `mapstructure:"HTTP_WRITE_TIMEOUT"`
		MaxHeaderBytes int           `mapstructure:"HTTP_MAX_HEADER_BYTES"`
	}

	// PG -.
	PG struct {
		URL     string `mapstructure:"PG_URL" validate:"required"`
		PoolMax int    `mapstructure:"PG_POOL_MAX"`
	}
)

func Load() (*Config, error) {
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Set defaults
	viper.SetDefault("GIN_MODE", "debug")
	viper.SetDefault("HTTP_READ_TIMEOUT", "10s")
	viper.SetDefault("HTTP_WRITE_TIMEOUT", "10s")
	viper.SetDefault("HTTP_MAX_HEADER_BYTES", "1")
	viper.SetDefault("PG_POOL_MAX", 2)

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	validate := validator.New()
	if err := validate.Struct(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
