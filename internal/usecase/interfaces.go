package usecase

import (
	"context"

	"github.com/chub-es/go-link-shortener/internal/entity"
)

type (
	Link interface {
		SearchLink(c context.Context, shortURL string) (entity.Link, error)
		CreateLink(c context.Context, l entity.Link) (string, error)
	}

	LinkRepo interface {
		Insert(c context.Context, link entity.Link) (string, error)
		FindOne(c context.Context, columns string, args ...interface{}) (entity.Link, error)
		UpShowned(c context.Context, linkID int64) error
	}
)
