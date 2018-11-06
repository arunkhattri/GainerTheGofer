package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Printf("%d\ttype: %T\n", a, a)
	fmt.Printf("%d\ttype: %T\n", b, b)
	a = int(b) // conversion, value of type hotdog
	fmt.Printf("%d\ttype: %T\n", a, a)

}
