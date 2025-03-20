package book

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID          uuid.UUID `json:"id" validate:"required" gorm:"primaryKey"`
	Title       string    `json:"title" validate:"required" gorm:"not null"`
	Author      string    `json:"author" validate:"required" gorm:"not null"`
	ISBN        string    `json:"isbn" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	Publisher   string    `json:"publisher" validate:"required" gorm:"not null"`
	Published   time.Time `json:"published" validate:"required" gorm:"not null"`
	Pages       int       `json:"pages" gorm:"not null"`
	Cover       string    `json:"cover,omitempty" gorm:"not null"`
	Genre       string    `json:"genre" validate:"required" gorm:"not null"`
}
