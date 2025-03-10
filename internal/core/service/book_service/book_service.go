package book_service

import (
	"GoChallenge/internal/core/domain/entity/book"
	"GoChallenge/internal/infrastructure/adapters/repositories/book_repository"
	"errors"
)

func GetAllBooks() []book.Book {
	return book_repository.FindAll()
}

func GetBookById(id string) (book.Book, bool) {
	return book_repository.FindById(id)
}

func SaveBook(newBook book.Book) error {
	if book_repository.Create(newBook) {
		return nil
	}

	return errors.New("error while saving book")
}

func UpdateBook(id string, book book.Book) error {
	if book_repository.Update(id, book) {
		return nil
	}

	return errors.New("book does not exists")
}

func DeleteBook(id string) error {
	if book_repository.Delete(id) {
		return nil
	}

	return errors.New("book does not exists")
}
