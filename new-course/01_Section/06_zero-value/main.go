package main

import "fmt"

var x int    // var x of type int is initialized with 0 value
var y string // var y of type string is initialized with 0 value

func main() {
	fmt.Println(x)
	fmt.Println(y)

	x = 11214
	fmt.Println(x)

	y = "Khattri, Aaryaansh Arora"
	fmt.Println(y)
}
