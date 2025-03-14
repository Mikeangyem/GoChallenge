package book_repository

import (
	"GoChallenge/internal/core/domain/entity/book"
	"time"

	"github.com/google/uuid"
)

var books = []book.Book{
	{
		ID:          uuid.New(),
		Title:       "Pedro Paramo",
		Author:      "Juan Rulfo",
		ISBN:        "",
		Description: "",
		Publisher:   "",
		Published:   time.Now(),
		Pages:       100,
		Cover:       "",
		Genre:       "Realismo Mágico",
	},
}

func FindAll() []book.Book {
	return books
}

func FindById(id uuid.UUID) (book.Book, bool) {
	for _, book := range books {
		if book.ID == id {
			return book, true
		}
	}

	notFoundBook := book.Book{
		ID:          uuid.Nil,
		Title:       "",
		Author:      "",
		ISBN:        "",
		Description: "",
		Publisher:   "",
		Published:   time.Now(),
		Pages:       0,
		Cover:       "",
		Genre:       "",
	}

	return notFoundBook, false
}

func Create(newBook book.Book) bool {
	books = append(books, newBook)

	return true
}

func Update(id uuid.UUID, updatedBook book.Book) bool {
	for i, book := range books {
		if id == book.ID {
			books[i] = updatedBook
			return true
		}
	}

	return false
}

func Delete(id uuid.UUID) bool {
	for i, book := range books {
		if id == book.ID {
			books = append(books[:i], books[i+1:]...)
			return true
		}
	}

	return false
}
