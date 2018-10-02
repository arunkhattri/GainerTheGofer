package main

import "fmt"

func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true

	// print values

	// fmt.Printf("%v \n", a)
	// fmt.Printf("%v \n", b)
	// fmt.Printf("%v \n", c)
	// fmt.Printf("%v \n", d)

	// print value and types

	fmt.Printf("a: %v \t %T \n", a, a)
	fmt.Printf("b: %v \t %T \n", b, b)
	fmt.Printf("c: %v \t %T \n", c, c)
	fmt.Printf("d: %v \t %T \n", d, d)
}
