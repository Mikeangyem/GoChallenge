package router

import (
	book_handler "GoChallenge/internal/infrastructure/entrypoints/handler/book"

	"github.com/gin-gonic/gin"
)

func Serve(bookHandler *book_handler.BookHandler) {
	router := gin.Default()
	router.GET("/books", bookHandler.GetAll)
	router.GET("/books/:id", bookHandler.GetById)
	router.POST("/books", bookHandler.Create)
	router.PUT("/books/:id", bookHandler.Update)
	router.DELETE("/books/:id", bookHandler.Delete)

	router.Run("0.0.0.0:8080")
}
