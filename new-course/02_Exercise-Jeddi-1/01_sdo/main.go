package main

import "fmt"

func main() {
	a := 42
	b := "James Bond"
	c := true

	fmt.Println(a)
	fmt.Printf("%s\t%T\n", b, b)
	fmt.Println(a, b, c)
}
