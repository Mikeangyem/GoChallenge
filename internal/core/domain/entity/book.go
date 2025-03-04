package entity

import "time"

type book struct {
	id          string
	title       string
	author      string
	isbn        string
	description string
	publisher   string
	published   time.Time
	pages       int
	cover       string
	genre       string
}
