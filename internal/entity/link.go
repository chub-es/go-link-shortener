package entity

import (
	"time"
)

type Link struct {
	CreatedAt   time.Time `json:"created_at"`
	OriginalURL string    `json:"original_url"`
	ShortURL    string    `json:"short_url"`
}
