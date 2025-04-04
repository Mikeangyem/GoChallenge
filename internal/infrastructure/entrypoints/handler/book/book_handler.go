package book_handler

import (
	"GoChallenge/internal/core/dto"
	ports "GoChallenge/internal/core/port"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type BookHandler struct {
	service ports.BookServiceInterface
}

func InitBookHandler(service ports.BookServiceInterface) *BookHandler {
	return &BookHandler{
		service: service,
	}
}

func (handler *BookHandler) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, handler.service.GetAllBooks())
}

func (handler *BookHandler) GetById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid book"})
		return
	}

	book, wasFound := handler.service.GetBookById(id)

	if wasFound {
		c.JSON(http.StatusOK, book)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
	}
}

func (handler *BookHandler) Create(c *gin.Context) {
	var newBook dto.BookDTO

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	validate := validator.New()
	if err := validate.Struct(newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid book"})
		return
	}

	if err := handler.service.SaveBook(&newBook); err != nil {
		c.JSON(http.StatusInternalServerError, newBook)
	} else {
		c.JSON(http.StatusCreated, newBook)
	}
}

func (handler *BookHandler) Update(c *gin.Context) {
	var bookDto dto.BookDTO
	if err := c.BindJSON(&bookDto); err != nil {
		return
	}

	validate := validator.New()
	if err := validate.Struct(bookDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid book"})
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid book"})
		return
	}

	if err := handler.service.UpdateBook(id, &bookDto); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, bookDto)
	}
}

func (handler *BookHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid book"})
		return
	}

	if err := handler.service.DeleteBook(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "book was deleted successfully"})
	}
}
