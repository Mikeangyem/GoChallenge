package config

import (
	"GoChallenge/internal/core/domain/entity/book"
	"GoChallenge/internal/infrastructure/entrypoints/router"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() {
	db, err := dbconnect()
	if err != nil {
		return
	}

	db.AutoMigrate(&book.Book{})
	db.Create(&book.Book{
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
	})

	router.Serve()
}

func dbconnect() (*gorm.DB, error) {
	dsn := "host=postgres user=postgres password=postgres dbname=books sslmode=disable"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
