package book_repository

import (
	"GoChallenge/internal/core/domain/entity/book"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*var books = []book.Book{
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
}*/

type bookRepository struct {
	db *gorm.DB
}

func InitBookRepository(db *gorm.DB) *bookRepository {
	return &bookRepository{
		db: db,
	}
}

func (repo *bookRepository) FindAll() ([]book.Book, error) {
	var bookSlice []book.Book
	result := repo.db.Find(&bookSlice)

	if result.Error != nil {
		return nil, result.Error
	}

	return bookSlice, nil
}

func (repo *bookRepository) FindById(id uuid.UUID) (*book.Book, error) {
	var bookToFind book.Book = book.Book{ID: id}
	result := repo.db.First(&bookToFind)

	if result.Error != nil {
		return &book.Book{}, result.Error
	}

	return &bookToFind, nil
}

func (repo *bookRepository) Create(newBook *book.Book) error {
	result := repo.db.Create(&newBook)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *bookRepository) Update(updatedBook *book.Book) error {
	/*bookToFind := book.Book{ID: updatedBook.ID}
	if r := repo.db.First(&bookToFind); r.Error != nil {
		return r.Error
	}*/

	result := repo.db.Save(&updatedBook)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *bookRepository) Delete(id uuid.UUID) error {
	result := repo.db.Delete(&book.Book{ID: id})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("not found")
	}

	return nil
}
