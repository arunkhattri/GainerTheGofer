package main

import "fmt"

var x = 43 // package level scope

func main() {
	y := 45 // func level scope
	fmt.Println(x)
	fmt.Println(y)

}
