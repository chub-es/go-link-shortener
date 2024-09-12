package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/chub-es/go-link-shortener/config"
	v1 "github.com/chub-es/go-link-shortener/internal/controller/http/v1"
	"github.com/chub-es/go-link-shortener/internal/usecase"
	"github.com/chub-es/go-link-shortener/internal/usecase/repo"
	"github.com/chub-es/go-link-shortener/pkg/httpserver"
	"github.com/chub-es/go-link-shortener/pkg/logger"
	"github.com/chub-es/go-link-shortener/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.HTTP.Mode)

	// Repostiry
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()
	l.Info("Postgres successfully connected")

	// Init Usecases
	linkUsecase := usecase.New(repo.New(pg))

	// Http server
	handler := gin.New()
	v1.NewRouter(handler, l, linkUsecase)
	server := httpserver.New(
		handler,
		httpserver.Port("8080"),
		httpserver.ReadTimeout(cfg.HTTP.ReadTimeout),
		httpserver.WriteTimeout(cfg.HTTP.WriteTimeout),
		httpserver.MaxHeaderBytes(cfg.HTTP.MaxHeaderBytes),
	)
	l.Info("Server started on port :%s", cfg.HTTP.Port)

	// Waiting signal
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	select {
	case s := <-channel:
		l.Info("app - Run - signal: " + s.String())
	case err := <-server.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = server.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
