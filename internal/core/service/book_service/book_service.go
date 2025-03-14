package book_service

import (
	"GoChallenge/internal/core/domain/entity/book"
	"GoChallenge/internal/core/dto"
	"GoChallenge/internal/infrastructure/adapters/repositories/book_repository"
	"errors"

	"github.com/google/uuid"
)

func GetAllBooks() []book.Book {
	return book_repository.FindAll()
}

func GetBookById(id uuid.UUID) (book.Book, bool) {
	return book_repository.FindById(id)
}

func SaveBook(bookDTO dto.BookDTO) error {
	newBook := book.Book{
		ID:          uuid.New(),
		Title:       bookDTO.Title,
		Author:      bookDTO.Author,
		ISBN:        bookDTO.ISBN,
		Description: bookDTO.Description,
		Publisher:   bookDTO.Publisher,
		Published:   bookDTO.Published,
		Pages:       bookDTO.Pages,
		Cover:       bookDTO.Cover,
		Genre:       bookDTO.Genre,
	}

	if book_repository.Create(newBook) {
		return nil
	}

	return errors.New("error while saving book")
}

func UpdateBook(id uuid.UUID, book book.Book) error {
	if book_repository.Update(id, book) {
		return nil
	}

	return errors.New("book does not exists")
}

func DeleteBook(id uuid.UUID) error {
	if book_repository.Delete(id) {
		return nil
	}

	return errors.New("book does not exists")
}
