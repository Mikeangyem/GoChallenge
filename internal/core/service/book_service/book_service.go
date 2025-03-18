package book_service

import (
	"GoChallenge/internal/core/domain/entity/book"
	"GoChallenge/internal/core/dto"
	ports "GoChallenge/internal/core/port"
	"errors"

	"github.com/google/uuid"
)

type bookService struct {
	repository ports.BookRepositoryInterface
}

func InitBookService(repository ports.BookRepositoryInterface) *bookService {
	return &bookService{
		repository: repository,
	}
}

func (service *bookService) GetAllBooks() []book.Book {
	books, err := service.repository.FindAll()

	if err != nil {
		return nil
	}

	return books
}

func (service *bookService) GetBookById(id uuid.UUID) (*book.Book, bool) {
	b, err := service.repository.FindById(id)

	if err != nil {
		return &book.Book{}, false
	}

	return b, true
}

func (service *bookService) SaveBook(bookDTO *dto.BookDTO) error {
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

	if err := service.repository.Create(&newBook); err == nil {
		return nil
	}

	return errors.New("error while saving book")
}

func (service *bookService) UpdateBook(id uuid.UUID, bookDTO *dto.BookDTO) error {
	updatedBook := book.Book{
		ID:          id,
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

	if err := service.repository.Update(&updatedBook); err == nil {
		return nil
	}

	return errors.New("book does not exists")
}

func (service *bookService) DeleteBook(id uuid.UUID) error {
	if err := service.repository.Delete(id); err == nil {
		return nil
	}

	return errors.New("book does not exists")
}
