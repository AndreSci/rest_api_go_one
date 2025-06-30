package models

type Book struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

type NewBook struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}
