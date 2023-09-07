package main

import (
	"syscall/js"
)

func addTwoNumbers(this js.Value, args []js.Value) interface{} {
	a := 2
	b := 4
	sum := a + b
	return js.ValueOf(sum)
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("addTwoNumbers", js.FuncOf(addTwoNumbers))
	<-c
}
