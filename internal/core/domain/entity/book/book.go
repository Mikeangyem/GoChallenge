package book

import "time"

type Book struct {
	ID          string
	Title       string
	Author      string
	ISBN        string
	Description string
	Publisher   string
	Published   time.Time
	Pages       int
	Cover       string
	Genre       string
}
