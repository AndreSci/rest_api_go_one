package server

import (
	"errors"
	"sync"
	"time"
)

// cache
var (
	books         = []Book{}
	timeForUpdate = 100
	timeUpdate    = time.Now().Add(-100 * time.Second)
	mu            sync.Mutex
)

type Book struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

type NewBook struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

func SearchBookByID(id int) (Book, int, error) {

	for i := range books {
		if books[i].Id == id {
			return books[i], i, nil
		}
	}

	return Book{-1, "", ""}, -1, errors.New("can't found data")
}
