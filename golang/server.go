package main

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/rs/cors"

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

type ReturnValue struct {
	ID int `json:"id"`
	Title string `json:"title"`
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
        json.NewEncoder(w).Encode(ReturnValue{
            Title: "Internal Server Error",
        })
        return
    }

	// var poems []Poem
	
	// db.Table("poems").Select("id, title").Find(&poems)
	
	// poems[0] = Poem{
	// 	ID: 1, 
	// 	Title: "Static data",
	// }

	var myPoems []Poem
	myPoems = append(myPoems, Poem{ID: 1, Title: "Static Data"})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(myPoems)
}

func NewConnection() (*gorm.DB, error) {
    configurations := Config{
        Host:     "localhost",
        Port:     "5432",
        Password: "Javohirjavohir1?",
        User:     "postgres",
        DBName:   "users",
        SSLMode:  "disable",
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