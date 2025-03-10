package router

import (
	book_handler "GoChallenge/internal/infrastructure/entrypoints/handler/book"

	"github.com/gin-gonic/gin"
)

func Serve() {
	router := gin.Default()
	router.GET("/books", book_handler.GetAll)
	router.GET("/books/:id", book_handler.GetById)
	router.POST("/books", book_handler.Create)
	router.PUT("/books/:id", book_handler.Update)
	router.DELETE("/books/:id", book_handler.Delete)

	router.Run("localhost:8080")
}
