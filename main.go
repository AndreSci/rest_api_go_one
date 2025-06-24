package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AndreSci/rest_api_go_one/server"
)

func main() {
	fmt.Println("Hello REST API with Golang")

	http.HandleFunc("/books", server.LoggerMiddleware(server.HandlerBooksGet))
	http.HandleFunc("/book", server.LoggerMiddleware(server.HandleBook))

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
