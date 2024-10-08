package usecase

import (
	"context"

	"github.com/chub-es/go-link-shortener/internal/entity"
)

type (
	Link interface {
		GetURL(c context.Context, shortURL string) (string, error)
		Create(c context.Context, l entity.Link) (entity.Link, error)
	}

	LinkRepo interface {
		Insert(c context.Context, link entity.Link) (int, error)
		FindOne(c context.Context, columns string, args interface{}) (entity.Link, error)
	}
)
