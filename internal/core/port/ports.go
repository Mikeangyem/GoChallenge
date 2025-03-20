package ports

import (
	"GoChallenge/internal/core/domain/entity/book"
	"GoChallenge/internal/core/dto"

	"github.com/google/uuid"
)

type BookRepositoryInterface interface {
	FindAll() ([]book.Book, error)
	FindById(id uuid.UUID) (*book.Book, error)
	Create(newBook *book.Book) error
	Update(updatedBook *book.Book) error
	Delete(id uuid.UUID) error
}

type BookServiceInterface interface {
	GetAllBooks() []book.Book
	GetBookById(id uuid.UUID) (*dto.BookDTO, bool)
	SaveBook(bookDTO *dto.BookDTO) error
	UpdateBook(id uuid.UUID, bookDTO *dto.BookDTO) error
	DeleteBook(id uuid.UUID) error
}
