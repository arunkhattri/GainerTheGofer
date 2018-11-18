package main

import "fmt"

const (
	a int = 2 // typed constant
	b     = 4 // untyped constant
)

func main() {
	fmt.Printf("%d\t%T\n", a, a)
	fmt.Printf("%d\t%T", b, b)
}
