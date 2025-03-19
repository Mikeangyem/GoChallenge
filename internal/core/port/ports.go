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
	GetBookById(id uuid.UUID) (*book.Book, bool)
	SaveBook(bookDTO *dto.BookDTO) error
	UpdateBook(id uuid.UUID, bookDTO *dto.BookDTO) error
	DeleteBook(id uuid.UUID) error
}

/*type BookHandlerInterface interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}*/
