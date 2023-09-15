package main

import (
	// "fmt"
	"log"
	"net/http"
	"syscall/js"

	// "github.com/rs/cors"

	// "github.com/golang-jwt/jwt"
	// "golang.org/x/crypto/bcrypt"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
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

func addTwoNumbers(this js.Value, args []js.Value) interface{} {
	a := args[0].Int()
	b := args[1].Int()
	sum := a + b
	return js.ValueOf(sum)
}

func getPoems(this js.Value, args []js.Value) interface{} {
	url := "http://localhost:9090/poems"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
		
	return js.ValueOf(resp.Body)
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("addTwoNumbers", js.FuncOf(addTwoNumbers))
	js.Global().Set("getPoems", js.FuncOf(getPoems))
	<-c
}
