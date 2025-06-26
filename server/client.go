package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/AndreSci/rest_api_go_one/pkg"
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
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}

func (c Client) GetBooks() ([]byte, error) {

	mu.Lock()
	defer mu.Unlock()

	currentTime := time.Now()
	timeSpendAfter := int(currentTime.Sub(timeUpdate).Seconds())

	if timeForUpdate < timeSpendAfter {
		fmt.Println("Try to update books")
		// TODO SQL REQUEST HERE
		err := selectBooksPostgres()

		if err != nil {
			return nil, err
		}
		timeUpdate = time.Now()
	}

	fmt.Println(books)

	if len(books) < 1 {
		return nil, errors.New("something went wrong: no data or connection to the database")
	}

	return json.Marshal(books)
}

func (c Client) GetBookById(id int) ([]byte, error) {
	mu.Lock()
	defer mu.Unlock()

	currentTime := time.Now()
	timeSpendAfter := int(currentTime.Sub(timeUpdate).Seconds())

	if timeForUpdate < timeSpendAfter {
		fmt.Println("Try to update books")
		// TODO SQL REQUEST HERE
		//books = genBooks()
		err := selectBooksPostgres()

		if err != nil {
			return nil, err
		}

		timeUpdate = time.Now()
	}

	book, _, err := SearchBookByID(id)
	if err != nil {
		return nil, err
	}

	return json.Marshal(book)
}

func (c Client) AddBook(newBook *NewBook) error {

	tx, err := pkg.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("insert into books (name, author) values ($1, $2);", newBook.Name, newBook.Author)

	if err != nil {
		return err
	}
	fmt.Println("Success add book")
	timeUpdate = time.Now().Add(-100 * time.Second)

	return tx.Commit()
}

func (c Client) DeleteBook(id int) error {
	tx, err := pkg.DB.Begin()

	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("delete from books where id = $1;", id)

	if err != nil {
		return err
	}
	fmt.Println("Success delete book")
	timeUpdate = time.Now().Add(-100 * time.Second)

	return tx.Commit()
}

func (c Client) UpdateBook(updateBook *Book) error {

	_, index, err := SearchBookByID(updateBook.Id)

	if err != nil {
		return err
	}
	books[index] = *updateBook

	return nil
}

// Жестко но что поделать :)
func selectBooksPostgres() error {

	rows, err := pkg.DB.Query("select * from books;")
	if err != nil {
		return err
	}

	booksNew := make([]Book, 0)

	for rows.Next() {
		b := Book{}
		err := rows.Scan(&b.Id, &b.Name, &b.Author)
		if err != nil {
			return err
		}

		booksNew = append(booksNew, b)
	}

	books = booksNew

	return nil
}

func (c Client) DeleteAll() error {
	tx, err := pkg.DB.Begin()

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
	fmt.Println("Success DROP TABLE books")
	timeUpdate = time.Now().Add(-100 * time.Second)

	return tx.Commit()
}
