package usecase

import (
	"context"

	"github.com/chub-es/go-link-shortener/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	Link interface {
		SearchLink(c context.Context, l entity.Link) (entity.Link, error)
		CreateLink(c context.Context, l entity.Link) (string, error)
		UpShownLink(c context.Context, l entity.Link) error
	}

	LinkRepo interface {
		Insert(c context.Context, link entity.Link) (string, error)
		FindOne(c context.Context, columns string, args ...any) (entity.Link, error)
		UpShowned(c context.Context, link entity.Link) error
	}
)
