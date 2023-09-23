package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	// "io/ioutil"
	// "net/http"
	"syscall/js"

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
	ID    uint   `json:"ID"`
	Title string `json:"Title"`
}

// type Poem struct {
// 	ID    uint
// 	Title string
// }

func addTwoNumbers(this js.Value, args []js.Value) interface{} {
	a := args[0].Int()
	b := args[1].Int()
	sum := a + b
	return js.ValueOf(sum)
}

func getPoemss(this js.Value, inputs []js.Value) interface{} {
	fmt.Println("I am here!")
	db, err := NewConnection()
	if err != nil {
		fmt.Println("Errorr:", err)
		return nil
	}
	defer CloseConnection(db)

	var poems []Poem
	db.Find(&poems)
	fmt.Println("PEOMS JSON", poems)

    // Convert the Go slice to a JavaScript array
	poemsJSON, err := json.Marshal(poems)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	returningValue := js.ValueOf(string(poemsJSON))
	fmt.Println("RETURNING VALUE", returningValue)
	fmt.Println("POEMSJSON", string(poemsJSON))
	return returningValue
	// return js.Global().Get("JSON").Call("parse", returningValue)
}

func NewConnection() (*gorm.DB, error) {
	fmt.Println("I am in NEWCOnnection")
    configurations := Config{
		Host:     "localhost",
        Port:     "5432",
        Password: "Javohirjavohir1?",
        User:     "postgres",
        DBName:   "users",
        SSLMode:  "disable",
    }
	
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", configurations.Host, configurations.Port, configurations.User, configurations.Password, configurations.DBName, configurations.SSLMode)
	fmt.Println("DSN", dsn)
	post := postgres.New(postgres.Config{DSN: dsn})
	fmt.Println("POST", post)
	db, err := gorm.Open(post, &gorm.Config{})
	fmt.Println("DB", db)
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

func getPoems(this js.Value, inputs []js.Value) interface{} {
	request, err := http.NewRequest("GET", "http://localhost:9000/poems", nil)
	// request, err := http.Get("http://localhost:9000/poems")
	if err != nil {
		fmt.Println("ERROR", err)
		return nil
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("ERROR 111", err)
		return nil
	}

	defer response.Body.Close()

	// url := "http://localhost:9000/poems"
	// response, err := http.Get(url)
	// if err != nil {
	// 	fmt.Println("Errorr:", err)
	// 	return nil
	// }
	// defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	var poems []map[string]interface{}
	err = json.Unmarshal(body, &poems)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	fmt.Println("AAAPOEMSAAA", poems)
	return js.ValueOf(poems)
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("addTwoNumbers", js.FuncOf(addTwoNumbers))
	js.Global().Set("getPoems", js.FuncOf(getPoems))
	<-c
}
