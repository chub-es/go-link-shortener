package entity

import (
	"time"
)

type Link struct {
	ID          int64     `db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	OriginalURL string    `json:"original_url" db:"original_url"`
	ShortURL    string    `json:"short_url" db:"short_url"`
	Showned int64 `db:"showned"`
}
