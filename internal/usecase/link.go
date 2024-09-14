package usecase

import (
	"context"
	"errors"

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

// GetURL -.
func (uc *LinkUseCase) GetURL(c context.Context, shortURL string) (string, error) {
	link, err := uc.repo.FindOne(c, "short_url = ?", shortURL)
	if err != nil {
		return "", err
	}

	if link.OriginalURL == "" {
		return "", errors.New("unknown link")
	}

	return link.OriginalURL, nil
}

// Create -.
func (uc *LinkUseCase) Create(c context.Context, l entity.Link) (string, error) {
	shortURL, err := uc.repo.Insert(c, l)
	if err != nil {
		return "", err
	}

	// link, err := uc.repo.FindOne(c, "id = ?", ID)
	// if err != nil {
	// 	return entity.Link{}, err
	// }

	return shortURL, nil
}
