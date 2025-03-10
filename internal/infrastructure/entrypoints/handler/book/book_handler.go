package book_handler

import (
	"GoChallenge/internal/core/domain/entity/book"
	"GoChallenge/internal/core/service/book_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, book_service.GetAllBooks())
}

func GetById(c *gin.Context) {
	id := c.Param("id")
	book, wasFound := book_service.GetBookById(id)

	if wasFound {
		c.IndentedJSON(http.StatusOK, book)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
	}
}

func Create(c *gin.Context) {
	var newBook book.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	if err := book_service.SaveBook(newBook); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, newBook)
	} else {
		c.IndentedJSON(http.StatusCreated, newBook)
	}
}

func Update(c *gin.Context) {
	var book book.Book
	if err := c.BindJSON(&book); err != nil {
		return
	}

	id := c.Param("id")
	if err := book_service.UpdateBook(id, book); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
	} else {
		c.IndentedJSON(http.StatusOK, book)
	}
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	if err := book_service.DeleteBook(id); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "book was deleted successfully"})
	}
}
