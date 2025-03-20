package config

import (
	"GoChallenge/internal/core/domain/entity/book"
	"GoChallenge/internal/core/service/book_service"
	"GoChallenge/internal/infrastructure/adapters/repositories/book_repository"
	book_handler "GoChallenge/internal/infrastructure/entrypoints/handler/book"
	"GoChallenge/internal/infrastructure/entrypoints/router"
	"fmt"

	"gorm.io/driver/mysql"
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
	// Postgress/Docker
	/*
		dsn := "host=postgres user=postgres password=postgres dbname=books sslmode=disable"
		return gorm.Open(postgres.Open(dsn), &gorm.Config{})
	*/

	// MySQL
	//dbDriver := "mysql"
	dbName := "books"
	dbUser := "mysql"
	dbPassword := "mysql"
	dbHost := "127.0.0.1"
	dbPort := "3306"
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	return gorm.Open(mysql.Open(connectionString), &gorm.Config{})
}
