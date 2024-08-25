package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/chub-es/go-link-shortener/config"
	"github.com/chub-es/go-link-shortener/pkg/httpserver"
	"github.com/chub-es/go-link-shortener/pkg/logger"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Level)

	h := gin.New()
	server := httpserver.New(h, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-server.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err := server.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
