package main

// go get -u github.com/lib/pq
// требуется для установки драйвера коннекта к БД postgres
// _ "github.com/lib/pq" // Имптор для сторонних эффектов

import (
	"database/sql"
	"fmt"

	"net/http"
	"os"

	"github.com/AndreSci/rest_api_go_one/internal/config"
	"github.com/AndreSci/rest_api_go_one/internal/models"
	"github.com/AndreSci/rest_api_go_one/internal/service"
	"github.com/AndreSci/rest_api_go_one/pkg"
	unittest_test "github.com/AndreSci/rest_api_go_one/unit-tests"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq" // Имптор для сторонних эффектов

	log "github.com/sirupsen/logrus"
)

const (
	CONFIG_DIR  = "configs"
	CONFIG_FILE = "main"
)

// Обьявляет логер Logrus
func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

// @title Books APP
// @version 1.0
// @description API Server for Books

// @host 127.0.0.1:8081
// @BasePath /

func main() {
	_ = godotenv.Load(".env") // загружает переменные из файла .env ДЛЯ WINDOWS
	//fmt.Println("Hello REST API with Golang")
	log.Info("Hello REST API with Golang")

	cfg, err := config.New(CONFIG_DIR, CONFIG_FILE)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("config: %+v\n", *cfg)

	// CONNECT TO DB
	// NEED to import POSTGRES driver
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Name,
		cfg.DB.SSLMode)

	models.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Exception in:", err)
	}
	defer models.DB.Close()

	if err = models.DB.Ping(); err != nil {
		log.Fatal("No connection to the DataBase:", err)
	}

	// CREATE HANDLES FUNC
	http.HandleFunc("/books", pkg.LoggerMiddleware(service.HandlerBooksGet))
	http.HandleFunc("/book", pkg.LoggerMiddleware(service.HandleBook))

	// AUTO-TEST
	go unittest_test.RunTests()

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), nil); err != nil {
		log.Fatal(err)
	}
}
