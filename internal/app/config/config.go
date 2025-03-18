package config

import (
	"GoChallenge/internal/core/domain/entity/book"
	"GoChallenge/internal/core/service/book_service"
	"GoChallenge/internal/infrastructure/adapters/repositories/book_repository"
	book_handler "GoChallenge/internal/infrastructure/entrypoints/handler/book"
	"GoChallenge/internal/infrastructure/entrypoints/router"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() {
	db, err := dbconnect()
	if err != nil {
		return
	}

	db.AutoMigrate(&book.Book{})
	/*db.Create(&book.Book{
		ID:          uuid.New(),
		Title:       "Pedro Paramo",
		Author:      "Juan Rulfo",
		ISBN:        "",
		Description: "",
		Publisher:   "Fundación Juan Rulfo",
		Published:   time.Now(),
		Pages:       100,
		Cover:       "",
		Genre:       "Realismo Mágico",
	})*/

	bookRepository := book_repository.InitBookRepository(db)
	bookService := book_service.InitBookService(bookRepository)
	bookHandler := book_handler.InitBookHandler(bookService)

	router.Serve(bookHandler)
}

func dbconnect() (*gorm.DB, error) {
	dsn := "host=postgres user=postgres password=postgres dbname=books sslmode=disable"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
