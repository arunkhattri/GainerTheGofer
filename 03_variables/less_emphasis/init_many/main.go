package main

import "fmt"

func main() {
	var message = "Hello World"
	var a, b, c int = 1, 2, 3

	// can also be done this inside a func
	mesg := "Hello Arun"
	A, B, C := 1, false, 3
	d := 4
	e := true

	fmt.Println(message, a, b, c)
	fmt.Println(mesg, A, B, C, d, e)
}
