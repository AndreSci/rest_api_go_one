package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
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
		books = []Book{
			{1, "Book1", "Author1"},
			{2, "Book2", "Author2"},
			{3, "Book3", "Author3"},
			{4, "Book4", "Author4"},
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
		books = []Book{
			{1, "Book1", "Author1"},
			{2, "Book2", "Author2"},
			{3, "Book3", "Author3"},
			{4, "Book4", "Author4"},
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

	// TODO SQL REQUEST
	lastId := books[len(books)-1].Id

	neme := newBook.Name
	author := newBook.Author

	newBookReb := Book{lastId + 1, neme, author}

	books = append(books, newBookReb)

	return nil
}

func (c Client) UpdateBook(updateBook *Book) error {

	_, index, err := SearchBookByID(updateBook.Id)

	if err != nil {
		return err
	}
	books[index] = *updateBook

	return nil
}
