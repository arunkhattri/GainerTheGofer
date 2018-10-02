package main

import "fmt"

// accessible by all, package level scope
var x int = 42

func main() {
	fmt.Println(x)
	foo()
	// function level scope
	y := 17
	fmt.Println(y)
}

func foo() {
	fmt.Println(x)
}
