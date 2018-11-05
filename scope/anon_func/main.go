package main

import "fmt"

func main() {
	x := 0
	// anonymous function, function without a name
	// assigning a func to a variable (func expression)
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

// func increment(x int) int {
//     x++
//     return x
// }
//
// func main() {
//     fmt.Println(increment(1))
//     fmt.Println(increment(2))
// }
