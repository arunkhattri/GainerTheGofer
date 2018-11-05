package main

import "fmt"

// returning function
func wrapper() func() int { // outer
	x := 0
	return func() int { // inner
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}
