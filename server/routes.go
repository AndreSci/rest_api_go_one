package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func HandlerBooksGet(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	client, err := NewClient(time.Second * 10)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	byteBooks, err := client.GetBooks()

	if err != nil {
		fmt.Println(err)
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

func handlerBookByIdGet(w http.ResponseWriter, r *http.Request) {
	client, err := NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
		w.Write([]byte(err.Error()))
		return
	}

	idFromCx := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idFromCx)

	if err != nil {
		fmt.Printf("convertion error: Wrong type of ID in Body: %d\n", id)
		w.Write([]byte("convertion error: Wrong type of ID in Body"))
		return
	}

	byteBook, err := client.GetBookById(id)

	if err != nil {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(byteBook)
}

func handlerBookByIdUpdate(w http.ResponseWriter, r *http.Request) {
	client, err := NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
		w.Write([]byte(err.Error()))
		return
	}

	var book Book
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		fmt.Printf("Error read Body: %d\n", err)
		w.Write([]byte("Error Update book"))
		return
	}

	err = client.UpdateBook(&book)

	if err != nil {
		fmt.Printf("Error with add book: %d\n", err)
		w.Write([]byte("Error Update book"))
		return
	}

	w.Write([]byte("Success Update book"))
}

func handlerBookAdd(w http.ResponseWriter, r *http.Request) {
	client, err := NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
		w.Write([]byte(err.Error()))
		return
	}

	var book NewBook
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		fmt.Printf("Error read Body: %d\n", err)
	}

	err = client.AddBook(&book)

	if err != nil {
		fmt.Printf("Error with add book: %d\n", err)
	}

	w.Write([]byte("Success add book"))
}

func handlerBookDelete(w http.ResponseWriter, r *http.Request) {
	client, err := NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
		w.Write([]byte(err.Error()))
		return
	}

	idFromCx := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idFromCx)

	if err != nil {
		fmt.Printf("convertion error: Wrong type of ID in Body: %d\n", id)
		w.Write([]byte("convertion error: Wrong type of ID in Body"))
		return
	}

	err = client.DeleteBook(id)

	if err != nil {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Success delete book"))
}
