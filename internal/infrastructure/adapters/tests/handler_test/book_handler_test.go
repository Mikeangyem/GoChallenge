package handler_test

import (
	"GoChallenge/internal/core/domain/entity/book"
	"GoChallenge/internal/core/dto"
	"GoChallenge/internal/core/service/book_service"
	"GoChallenge/internal/infrastructure/adapters/repositories/book_repository"
	"GoChallenge/internal/infrastructure/adapters/tests"
	book_handler "GoChallenge/internal/infrastructure/entrypoints/handler/book"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var dB *gorm.DB
var bookHandler *book_handler.BookHandler
var router *gin.Engine

var bookDtos = []dto.BookDTO{
	{
		Title:       "Pedro Paramo",
		Author:      "Juan Rulfo",
		ISBN:        "9789685208550",
		Description: "",
		Publisher:   "Fondo de Cultura Economica",
		Published:   time.Date(1956, 1, 3, 0, 0, 0, 0, time.Local),
		Pages:       132,
		Cover:       "",
		Genre:       "Realismo Mágico",
	},
	{
		Title:       "El llano en llamas",
		Author:      "Juan Rulfo",
		ISBN:        "9788493442613",
		Description: "",
		Publisher:   "RM",
		Published:   time.Date(1953, 9, 6, 0, 0, 0, 0, time.Local),
		Pages:       169,
		Cover:       "",
		Genre:       "Realismo Mágico",
	},
	{
		Title:       "Habitos Atómicos",
		Author:      "James Clear",
		ISBN:        "978968529870",
		Description: "",
		Publisher:   "Paidos",
		Published:   time.Date(2010, 9, 6, 0, 0, 0, 0, time.Local),
		Pages:       350,
		Cover:       "",
		Genre:       "Otro",
	},
}

func TestMain(m *testing.M) {
	db, err := tests.SetUpTestDB()
	if err != nil {
		log.Fatalf("Error setting up test database: %v", err)
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book_repository.InitBookRepository(db)
	bookService := book_service.InitBookService(bookRepository)
	bookHandler = book_handler.InitBookHandler(bookService)

	dB = db

	router = tests.SetUpTestRouter()

	code := m.Run()

	tests.EmptyDB(db)

	os.Exit(code)
}

func TestCreate(t *testing.T) {
	router.POST("/books", bookHandler.Create)

	for _, bookDto := range bookDtos {
		w := httptest.NewRecorder()

		bookJson, _ := json.Marshal(bookDto)
		req, _ := http.NewRequest("POST", "/books", strings.NewReader(string(bookJson)))
		router.ServeHTTP(w, req)

		assert.Equal(t, 201, w.Code)
		assert.Equal(t, string(bookJson), w.Body.String())
	}
}

func TestGetAllBooks(t *testing.T) {
	router.GET("/books", bookHandler.GetAll)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/books", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var booksSlice []book.Book
	if err := json.Unmarshal(w.Body.Bytes(), &booksSlice); err != nil {
		log.Fatalf("Error converting JSON: %v", err)
	}
	assert.Equal(t, 3, len(booksSlice))
}

func TestGetById(t *testing.T) {
	router.GET("/books/:id", bookHandler.GetById)

	for _, bookDto := range bookDtos {
		var book book.Book
		result := dB.Where(
			"title = ? AND author = ? AND publisher = ?",
			bookDto.Title,
			bookDto.Author,
			bookDto.Publisher,
		).First(&book)
		if result.Error != nil {
			log.Fatalf("Error retrieving data from DB: %v", result.Error)
		}

		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", "/books/"+book.ID.String(), nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		bookJson, _ := json.Marshal(bookDto)
		assert.Equal(t, string(bookJson), w.Body.String())
	}
}

func TestUpdate(t *testing.T) {
	router.PUT("/books/:id", bookHandler.Update)

	for _, bookDto := range bookDtos {
		var book book.Book
		result := dB.Where(
			"title = ? AND author = ? AND publisher = ?",
			bookDto.Title,
			bookDto.Author,
			bookDto.Publisher,
		).First(&book)
		if result.Error != nil {
			log.Fatalf("Error retrieving data from DB: %v", result.Error)
		}

		bookDto.Description = "Esta es una descripción genérica..."
		bookDto.Cover = "Esta caratula no existe..."

		w := httptest.NewRecorder()

		bookJson, _ := json.Marshal(bookDto)
		req, _ := http.NewRequest("PUT", "/books/"+book.ID.String(), strings.NewReader(string(bookJson)))
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, string(bookJson), w.Body.String())

		w2 := httptest.NewRecorder()
		req2, _ := http.NewRequest("GET", "/books/"+book.ID.String(), nil)
		router.ServeHTTP(w2, req2)

		assert.Equal(t, 200, w2.Code)
		assert.Equal(t, string(bookJson), w2.Body.String())
	}
}

func TestDelete(t *testing.T) {
	router.DELETE("/books/:id", bookHandler.Delete)

	for _, bookDto := range bookDtos {
		var book book.Book
		result := dB.Where(
			"title = ? AND author = ? AND publisher = ?",
			bookDto.Title,
			bookDto.Author,
			bookDto.Publisher,
		).First(&book)
		if result.Error != nil {
			log.Fatalf("Error retrieving data from DB: %v", result.Error)
		}

		w := httptest.NewRecorder()

		req, _ := http.NewRequest("DELETE", "/books/"+book.ID.String(), nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		w2 := httptest.NewRecorder()
		req2, _ := http.NewRequest("GET", "/books/"+book.ID.String(), nil)
		router.ServeHTTP(w2, req2)

		assert.Equal(t, 404, w2.Code)
	}
}

/*
func areBooksEqual(book1 book.Book, book2 book.Book) bool {
	return reflect.DeepEqual(book1, book2)
}*/
