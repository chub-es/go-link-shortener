package usecase

import (
	"context"

	"github.com/chub-es/go-link-shortener/internal/entity"
)

// LinkUseCase -.
type LinkUseCase struct {
	repo LinkRepo
}

// New -.
func New(r LinkRepo) *LinkUseCase {
	return &LinkUseCase{r}
}

// SearchLink -.
func (uc *LinkUseCase) SearchLink(c context.Context, shortURL string) (entity.Link, error) {
	link, err := uc.repo.FindOne(c, "short_url = ?", shortURL)
	if err != nil {
		return entity.Link{}, err
	}
	_ = uc.repo.UpShowned(c, link.ID)

	return link, nil
}

// Create -.
func (uc *LinkUseCase) CreateLink(c context.Context, l entity.Link) (string, error) {
	shortURL, err := uc.repo.Insert(c, l)
	if err != nil {
		return "", err
	}

	return shortURL, nil
}
