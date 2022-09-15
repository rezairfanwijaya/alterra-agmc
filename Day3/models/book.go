package models

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Page   int    `json:"page"`
	Writer string `json:"writer"`
}
