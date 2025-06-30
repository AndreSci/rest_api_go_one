package main

// go get -u github.com/lib/pq
// требуется для установки драйвера коннекта к БД postgres
// _ "github.com/lib/pq" // Имптор для сторонних эффектов

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/AndreSci/rest_api_go_one/internal/config"
	"github.com/AndreSci/rest_api_go_one/pkg"
	"github.com/AndreSci/rest_api_go_one/server"
	unittest_test "github.com/AndreSci/rest_api_go_one/unit-tests"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq" // Имптор для сторонних эффектов
)

const (
	CONFIG_DIR  = "configs"
	CONFIG_FILE = "main"
)

// @title Books APP
// @version 1.0
// @description API Server for Books

// @host 127.0.0.1:8081
// @BasePath /

func main() {
	_ = godotenv.Load(".env") // загружает переменные из файла .env ДЛЯ WINDOWS
	fmt.Println("Hello REST API with Golang")

	cfg, err := config.New(CONFIG_DIR, CONFIG_FILE)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("config: %+v\n", cfg)

	// CONNECT TO DB
	// NEED to import POSTGRES driver
	connStr := "host=127.0.0.1 port=5432 user=postgres password=goLANGn1nja dbname=postgres sslmode=disable"

	pkg.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Exception in:", err)
	}
	defer pkg.DB.Close()

	if err = pkg.DB.Ping(); err != nil {
		log.Fatal("No connection to the DataBase:", err)
	}

	// CREATE HANDLES FUNC
	http.HandleFunc("/books", server.LoggerMiddleware(server.HandlerBooksGet))
	http.HandleFunc("/book", server.LoggerMiddleware(server.HandleBook))

	// AUTO-TEST
	go unittest_test.RunTests()

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), nil); err != nil {
		log.Fatal(err)
	}
}
