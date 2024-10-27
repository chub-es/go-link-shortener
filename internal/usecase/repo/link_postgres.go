package repo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/chub-es/go-link-shortener/internal/entity"
	"github.com/chub-es/go-link-shortener/pkg/postgres"
)

// LinkRepo -.
type LinkRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *LinkRepo {
	return &LinkRepo{pg}
}

func (r *LinkRepo) Insert(c context.Context, link entity.Link) (string, error) {
	sql, args, err := r.Builder.
		Insert("links").
		Columns("original_url").
		Values(link.OriginalURL).
		Suffix("RETURNING \"short_url\"").
		ToSql()
	if err != nil {
		return "", fmt.Errorf("LinkRepo - Insert - r.Builder: %w", err)
	}

	var shortURL string
	row := r.Pool.QueryRow(c, sql, args...)
	if err = row.Scan(&shortURL); err != nil {
		return "", fmt.Errorf("LinkRepo - Insert - row.Scan: %w", err)
	}

	return shortURL, nil
}

func (r *LinkRepo) FindOne(c context.Context, columns string, args ...interface{}) (entity.Link, error) {
	sql, _, err := r.Builder.
		Select("*").
		From("links").
		Where(columns).
		OrderBy("created_at DESC").
		Limit(1).
		ToSql()
	if err != nil {
		return entity.Link{}, fmt.Errorf("LinkRepo - FindOne - r.Builder: %w", err)
	}

	var link entity.Link
	row := r.Pool.QueryRow(c, sql, args...)
	if err = row.Scan(&link.ID, &link.CreatedAt, &link.OriginalURL, &link.ShortURL, &link.Showned); err != nil {
		return entity.Link{}, fmt.Errorf("LinkRepo - FindOne - row.Scan: %w", err)
	}

	return link, nil
}

func (r *LinkRepo) UpShowned(c context.Context, linkID int64) error {
	sql, args, err := r.Builder.
		Update("links").
		Set("showned", squirrel.Expr("showned + 1")).
		Where(squirrel.Eq{"id": linkID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("LinkRepo - UpShowned - r.Builder: %w", err)
	}

	_, err = r.Pool.Query(c, sql, args...)
	if err != nil {
		return fmt.Errorf("LinkRepo - UpShowned - r.Pool.Query: %w", err)
	}

	return nil
}
