package tests

import (
	"GoChallenge/internal/core/domain/entity/book"
	"reflect"
)

func TestGetAllBooks() {
	/*cases := []book.Book{
		{
			Title:       "Pedro Paramo",
			Author:      "Juan Rulfo",
			ISBN:        "9789685208550",
			Description: "",
			Publisher:   "Fondo de Cultura Economica",
			Published:   time.Date(1956, 1, 3, 0, 0, 0, 0, &time.Location{}),
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
			Published:   time.Date(1953, 9, 6, 0, 0, 0, 0, &time.Location{}),
			Pages:       169,
			Cover:       "",
			Genre:       "Realismo Mágico",
		},
	}*/
}

func areBooksEqual(book1 book.Book, book2 book.Book) bool {
	return reflect.DeepEqual(book1, book2)
}
