package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

func main() {
	fmt.Printf("%d, %T\n", a, a)
	fmt.Printf("%d, %T\n", b, b)
	fmt.Printf("%d, %T\n", c, c)
}
