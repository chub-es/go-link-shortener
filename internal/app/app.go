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

	// sql, args, _ := pg.Builder.
	// 	Select("*").
	// 	From("links").
	// 	Where("short_url = ?", "Fx2F2O").
	// 	OrderBy("created_at DESC").
	// 	Limit(1).
	// 	ToSql()
	// rows, _ := pg.Pool.Query(context.TODO(), sql, args...)
	// var link entity.Link
	// if rows.Next() {
	// 	rows.Scan(&link.ID, &link.CreatedAt, &link.OriginalURL, &link.ShortURL, &link.Showned)
	// }

	// sql, args, _ = pg.Builder.
	// 	Update("links").
	// 	Set("showned", squirrel.Expr("showned + 1")).
	// 	Where(squirrel.Eq{"id": link.ID}).
	// 	ToSql()
	// _, _ = pg.Pool.Query(context.TODO(), sql, args...)
	// log.Fatal("stop")

	// Init Usecases
	linkUsecase := usecase.New(repo.New(pg))

	// Http server
	handler := gin.New()
	v1.NewRouter(handler, l, linkUsecase)
	server := httpserver.New(
		handler,
		httpserver.Port(cfg.HTTP.Port),
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
