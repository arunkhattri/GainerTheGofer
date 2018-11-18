package main

import "fmt"

var x bool // by default 0 (false)

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)

	a := 7
	b := 42
	fmt.Println("a is equal to b: ", a == b)
	fmt.Println("a is not equal to b: ", a != b)
	fmt.Println("a is greater than b: ", a > b)
	fmt.Println("a is less than b: ", a < b)

}
