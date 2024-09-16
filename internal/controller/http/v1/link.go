package v1

import (
	"net/http"

	"github.com/chub-es/go-link-shortener/internal/entity"
	"github.com/chub-es/go-link-shortener/internal/usecase"
	"github.com/chub-es/go-link-shortener/pkg/logger"
	"github.com/gin-gonic/gin"
)

type linkRoutes struct {
	log logger.Interface
	uc  usecase.Link
}

func newLinkRoutes(handler *gin.Engine, log logger.Interface, link usecase.Link) {
	l := &linkRoutes{log, link}

	// Shorter route
	handler.GET("/:short_url", l.doRedirect)
	// Api route
	api := handler.Group("/api/v1")
	{
		api.POST("/url", l.doCreate)
	}
}

func (r *linkRoutes) doRedirect(c *gin.Context) {
	originalURL, err := r.uc.GetURL(c.Request.Context(), c.Param("short_url"))
	if err != nil {
		r.log.Error(err, "http - v1 - doRedirect - r.uc.GetURL")
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "unknown link"})

		return
	}

	c.Redirect(http.StatusMovedPermanently, originalURL)
}

type createRequest struct {
	OriginalURL string `json:"original_url" binding:"required"`
}

func (r *linkRoutes) doCreate(c *gin.Context) {
	var request createRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.log.Error(err, "http - v1 - doCreateShortUrl")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})

		return
	}

	shortURL, err := r.uc.Create(
		c.Request.Context(),
		entity.Link{
			OriginalURL: request.OriginalURL,
		},
	)
	if err != nil {
		r.log.Error(err, "http - v1 - doCreateShortUrl - r.uc.Create")
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
}
