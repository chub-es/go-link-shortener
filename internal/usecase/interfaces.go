package usecase

import (
	"context"

	"github.com/chub-es/go-link-shortener/internal/entity"
)

type (
	Link interface {
		GetURL(c context.Context, shortURL string) (string, error)
		Create(c context.Context, l entity.Link) (string, error)
	}

	LinkRepo interface {
		Insert(c context.Context, link entity.Link) (string, error)
		FindOne(c context.Context, columns string, args ...interface{}) (entity.Link, error)
		SetShowned(c context.Context, linkID int64) error
	}
)
