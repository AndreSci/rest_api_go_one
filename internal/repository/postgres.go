package repository

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/AndreSci/rest_api_go_one/internal/cache"
	"github.com/AndreSci/rest_api_go_one/internal/models"
	"github.com/AndreSci/rest_api_go_one/pkg"

	log "github.com/sirupsen/logrus"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be zero")
	}
	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &pkg.LoggingRoundTripper{
				Logger: os.Stdout,
				Next:   http.DefaultTransport,
			},
		},
	}, nil
}

func (c Client) GetBooks() ([]byte, error) {

	cache.Mu.Lock()
	defer cache.Mu.Unlock()

	currentTime := time.Now()
	timeSpendAfter := int(currentTime.Sub(cache.TimeUpdate).Seconds())

	if cache.TimeForUpdate < timeSpendAfter {
		log.Info("Try to update books")
		// TODO SQL REQUEST HERE
		err := selectBooksPostgres()

		if err != nil {
			return nil, err
		}
		cache.TimeUpdate = time.Now()
	}

	log.Info(cache.Books)

	if len(cache.Books) < 1 {
		return nil, errors.New("something went wrong: no data or connection to the database")
	}

	return json.Marshal(cache.Books)
}

func (c Client) GetBookById(id int) ([]byte, error) {
	cache.Mu.Lock()
	defer cache.Mu.Unlock()

	currentTime := time.Now()
	timeSpendAfter := int(currentTime.Sub(cache.TimeUpdate).Seconds())

	if cache.TimeForUpdate < timeSpendAfter {
		log.Info("Try to update books")
		// TODO SQL REQUEST HERE
		//books = genBooks()
		err := selectBooksPostgres()

		if err != nil {
			return nil, err
		}

		cache.TimeUpdate = time.Now()
	}

	book, _, err := cache.SearchBookByID(id)
	if err != nil {
		return nil, err
	}

	return json.Marshal(book)
}

func (c Client) AddBook(newBook *models.NewBook) error {

	tx, err := models.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("insert into books (name, author) values ($1, $2);", newBook.Name, newBook.Author)

	if err != nil {
		return err
	}
	log.Info("Success add book")
	cache.TimeUpdate = time.Now().Add(-100 * time.Second)

	return tx.Commit()
}

func (c Client) DeleteBook(id int) error {
	tx, err := models.DB.Begin()

	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("delete from books where id = $1;", id)

	if err != nil {
		return err
	}
	log.Info("Success delete book")
	cache.TimeUpdate = time.Now().Add(-100 * time.Second)

	return tx.Commit()
}

func (c Client) UpdateBook(updateBook *models.Book) error {

	_, index, err := cache.SearchBookByID(updateBook.Id)

	if err != nil {
		return err
	}
	cache.Books[index] = *updateBook

	return nil
}

// Жестко но что поделать :)
func selectBooksPostgres() error {

	rows, err := models.DB.Query("select * from books;")
	if err != nil {
		return err
	}

	booksNew := make([]models.Book, 0)

	for rows.Next() {
		b := models.Book{}
		err := rows.Scan(&b.Id, &b.Name, &b.Author)
		if err != nil {
			return err
		}

		booksNew = append(booksNew, b)
	}

	cache.Books = booksNew

	return nil
}

func (c Client) DeleteAll() error {
	tx, err := models.DB.Begin()

	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("DROP TABLE IF EXISTS books;")

	if err != nil {
		return err
	}
	_, err = tx.Exec("CREATE TABLE books (id SERIAL PRIMARY KEY, name VARCHAR(255), author VARCHAR(255));")

	if err != nil {
		return err
	}
	log.Info("Success DROP TABLE books")
	cache.TimeUpdate = time.Now().Add(-100 * time.Second)

	return tx.Commit()
}
