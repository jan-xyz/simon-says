package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello world!")
	val := js.Global().Get("html")
	fmt.Println(val)
	fmt.Println("End")
}
