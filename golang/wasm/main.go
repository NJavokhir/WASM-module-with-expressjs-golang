package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"syscall/js"
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

func addTwoNumbers(this js.Value, args []js.Value) interface{} {
	a := args[0].Int()
	b := args[1].Int()
	sum := a + b
	return js.ValueOf(sum)
}

func getPoems(this js.Value, inputs []js.Value) interface{} {
	response, err := http.Get("http://localhost:9090/poems")
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	var poems []Poem
	err = json.Unmarshal(body, &poems)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return poems
}


func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("addTwoNumbers", js.FuncOf(addTwoNumbers))
	js.Global().Set("getPoems", js.FuncOf(getPoems))
	<-c
}
