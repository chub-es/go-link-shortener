package repo

import (
	"context"
	"fmt"

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

func (r *LinkRepo) Insert(c context.Context, link entity.Link) (int, error) {
	sql, args, err := r.Builder.
		Insert("links").
		Columns("original_url").
		Values(link.OriginalURL).
		Suffix("RETURNING \"short_url\"").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("LinkRepo - Insert - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(c, sql, args...)

	var ID int
	if err = row.Scan(ID); err != nil {
		return 0, fmt.Errorf("LinkRepo - Insert - row.Scan: %w", err)
	}

	return ID, nil
}

func (r *LinkRepo) FindOne(c context.Context, columns string, args interface{}) (entity.Link, error) {
	sql, _, err := r.Builder.
		Select("created_at, original_url, short_url").
		From("links").
		Where(columns).
		Limit(1).
		ToSql()
	if err != nil {
		return entity.Link{}, fmt.Errorf("LinkRepo - FindOne - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(c, sql, args)
	if err != nil {
		return entity.Link{}, fmt.Errorf("LinkRepo - FindOne - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	var link entity.Link
	if rows.Next() {
		if err = rows.Scan(&link.CreatedAt, &link.OriginalURL, &link.ShortURL); err != nil {
			return entity.Link{}, fmt.Errorf("LinkRepo - FindOne - rows.Scan: %w", err)
		}
	}

	return link, nil
}
