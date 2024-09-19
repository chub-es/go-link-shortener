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
		api.POST("/url", l.doCreateLink)
	}
}

func (r *linkRoutes) doRedirect(c *gin.Context) {
	link, err := r.uc.SearchLink(c.Request.Context(), c.Param("short_url"))
	if err != nil {
		r.log.Error(err, "http - v1 - doRedirect - r.uc.SearchLink")
		c.AbortWithStatus(http.StatusNotFound)

		return
	}

	c.Redirect(http.StatusMovedPermanently, link.OriginalURL)
}

type createRequest struct {
	OriginalURL string `json:"original_url" binding:"required"`
}

func (r *linkRoutes) doCreateLink(c *gin.Context) {
	var request createRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.log.Error(err, "http - v1 - doCreateLink")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})

		return
	}

	shortURL, err := r.uc.CreateLink(
		c.Request.Context(),
		entity.Link{
			OriginalURL: request.OriginalURL,
		},
	)
	if err != nil {
		r.log.Error(err, "http - v1 - doCreateLink - r.uc.CreateLink")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error creating link"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
}
