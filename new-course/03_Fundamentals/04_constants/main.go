package main

import "fmt"

const (
	a int = 42
	b     = 42.78
	c     = "James Bond"
)

func main() {
	fmt.Println(a)
	fmt.Printf("%f, %T\n", b, b)
	fmt.Println(c)
}
