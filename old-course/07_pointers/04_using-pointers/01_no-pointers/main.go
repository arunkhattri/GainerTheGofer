package main

import "fmt"

func zero(z int) int {
	z = 0
	return z
}

func main() {
	x := 5
	var y = zero(x)
	fmt.Println(y)
	fmt.Println(zero(x))
}
