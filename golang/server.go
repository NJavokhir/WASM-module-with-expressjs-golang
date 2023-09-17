package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/rs/cors"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
    Host     string
    Port     string
    Password string
    User     string
    DBName   string
    SSLMode  string
}

type Poem struct {
	ID    uint
	Title string
}

func main() {
	http.HandleFunc("/poems", getPoems)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(http.DefaultServeMux)
	http.ListenAndServe(":9090", handler)
}

func getPoems(w http.ResponseWriter, r *http.Request) {
	db, err := NewConnection()
	defer CloseConnection(db)

	if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(Poem{
            Title: "Internal Server Error",
        })
        return
    }

	var poems []Poem
	db.Find(&poems)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(poems)
}

func NewConnection() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

    configurations := Config{
        Host:     os.Getenv("HOST"),
        Port:     os.Getenv("POSTGRES_PORT"),
        Password: os.Getenv("PASSWORD"),
        User:     os.Getenv("USER"),
        DBName:   os.Getenv("DBNAME"),
        SSLMode:  os.Getenv("SSLMODE"),
    }

    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", configurations.Host, configurations.Port, configurations.User, configurations.Password, configurations.DBName, configurations.SSLMode)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

    return db, nil
}

func CloseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from Database")
	}
	dbSQL.Close()
}