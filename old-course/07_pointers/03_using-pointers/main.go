package main

import (
	"fmt"
)

func main() {
	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // memory address of a

	var b = &a

	fmt.Println(b)  // memory address of pointer
	fmt.Println(*b) // value at that memory address

	*b = 42        // change the value to 42 at that address
	fmt.Println(a) // 42

	// this is useful, we can pass a memory address instead of a bunch of values
	// and then we can still change the value of whatever is stored at that memory address
	// this makes our programs more performant
	// we dont have to pass around large amounts of data
	// we only have to pass around addresses
	// when we pass a memory address, we are passing a value
}
