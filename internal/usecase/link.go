package usecase

import (
	"context"
	"fmt"

	"github.com/chub-es/go-link-shortener/internal/entity"
)

// LinkUsecase -.
type LinkUsecase struct {
	repo LinkRepo
}

// New -.
func New(r LinkRepo) *LinkUsecase {
	return &LinkUsecase{r}
}

// SearchLink -.
func (uc *LinkUsecase) SearchLink(c context.Context, l entity.Link) (entity.Link, error) {
	link, err := uc.repo.FindOne(c, "short_url = ?", l.ShortURL)
	if err != nil {
		return entity.Link{}, fmt.Errorf("LinkUsecase - SearchLink - uc.repo.FindOne: %w", err)
	}

	return link, nil
}

// Create -.
func (uc *LinkUsecase) CreateLink(c context.Context, l entity.Link) (string, error) {
	shortURL, err := uc.repo.Insert(c, l)
	if err != nil {
		return "", fmt.Errorf("LinkUsecase - CreateLink - uc.repo.Insert: %w", err)
	}
	return shortURL, nil
}

// UpShowned -.
func (uc *LinkUsecase) UpShownLink(c context.Context, l entity.Link) error {
	err := uc.repo.UpShowned(c, l)
	if err != nil {
		return fmt.Errorf("LinkUsecase - UpShownLink - uc.repo.UpShowned: %w", err)
	}

	return nil
}
