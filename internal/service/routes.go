package service

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/AndreSci/rest_api_go_one/internal/models"
	"github.com/AndreSci/rest_api_go_one/internal/repository"
)

// @Summary Books
// @Tags books
// @Description get full list of books
// @Accept  nones
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
func HandlerBooksGet(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	client, err := repository.NewClient(time.Second * 10)
	if err != nil {
		log.Error(err)
		w.Write([]byte(err.Error()))
		return
	}

	byteBooks, err := client.GetBooks()

	if err != nil {
		log.Error(err)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(byteBooks)
}

func HandleBook(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		handlerBookByIdGet(w, r)
	case http.MethodPost:
		handlerBookAdd(w, r)
	case http.MethodPut:
		handlerBookByIdUpdate(w, r)
	case http.MethodDelete:
		handlerBookDelete(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

// @Summary Book
// @Tags book
// @Description Get book by ID
// @Accept  int
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
func handlerBookByIdGet(w http.ResponseWriter, r *http.Request) {
	client, err := repository.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
		w.Write([]byte(err.Error()))
		return
	}

	idFromCx := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idFromCx)

	if err != nil {
		log.Printf("convertion error: Wrong type of ID in Body: %d\n", id)
		w.Write([]byte("convertion error: Wrong type of ID in Body"))
		return
	}

	byteBook, err := client.GetBookById(id)

	if err != nil {
		log.Error(err)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(byteBook)
}

// @Summary Book
// @Tags book
// @Description Update book by ID
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
func handlerBookByIdUpdate(w http.ResponseWriter, r *http.Request) {
	client, err := repository.NewClient(time.Second * 10)
	if err != nil {
		log.Error(err)
		w.Write([]byte(err.Error()))
		return
	}

	var book models.Book
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Error(err)
		w.Write([]byte("Error Update book"))
		return
	}

	err = client.UpdateBook(&book)

	if err != nil {
		log.Printf("Error with add book: %d\n", err)
		w.Write([]byte("Error Update book"))
		return
	}

	w.Write([]byte("Success Update book"))
}

// @Summary Book
// @Tags book
// @Description add book
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
func handlerBookAdd(w http.ResponseWriter, r *http.Request) {
	client, err := repository.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
		w.Write([]byte(err.Error()))
		return
	}

	var book models.NewBook
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Printf("Error read Body: %d\n", err)
	}

	err = client.AddBook(&book)

	if err != nil {
		log.Printf("Error with add book: %d\n", err)
	}

	w.Write([]byte("Success add book"))
}

// @Summary Book
// @Tags book
// @Description Delete from Db book by ID
// @Accept  int
// @Produce  none
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
func handlerBookDelete(w http.ResponseWriter, r *http.Request) {
	client, err := repository.NewClient(time.Second * 10)
	if err != nil {
		log.Error(err)
		w.Write([]byte(err.Error()))
		return
	}

	idFromCx := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idFromCx)

	if err != nil {
		log.Printf("convertion error: Wrong type of ID in Body: %d\n", id)
		w.Write([]byte("convertion error: Wrong type of ID in Body"))
		return
	}

	err = client.DeleteBook(id)

	if err != nil {
		log.Error(err)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Success delete book"))
}
