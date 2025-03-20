package dto

import "time"

type BookDTO struct {
	Title       string    `json:"title" validate:"required"`
	Author      string    `json:"author" validate:"required"`
	ISBN        string    `json:"isbn"`
	Description string    `json:"description"`
	Publisher   string    `json:"publisher" validate:"required"`
	Published   time.Time `json:"published" validate:"required"`
	Pages       int       `json:"pages"`
	Cover       string    `json:"cover,omitempty"`
	Genre       string    `json:"genre" validate:"required"`
}
