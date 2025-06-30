package cache

import (
	"errors"
	"sync"
	"time"

	"github.com/AndreSci/rest_api_go_one/internal/models"
)

// cache
var (
	Books         = []models.Book{}
	TimeForUpdate = 100
	TimeUpdate    = time.Now().Add(-100 * time.Second)
	Mu            sync.Mutex
)

func SearchBookByID(id int) (models.Book, int, error) {

	for i := range Books {
		if Books[i].Id == id {
			return Books[i], i, nil
		}
	}

	return models.Book{Id: -1, Name: "", Author: ""}, -1, errors.New("can't found data")
}
