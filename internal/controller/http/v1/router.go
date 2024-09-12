package v1

import (
	"net/http"

	"github.com/chub-es/go-link-shortener/internal/usecase"
	"github.com/chub-es/go-link-shortener/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter(r *gin.Engine, log logger.Interface, link usecase.Link) {
	// Options
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// K8s probe
	r.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Api route
	newLinkRoutes(r, log, link)
}
