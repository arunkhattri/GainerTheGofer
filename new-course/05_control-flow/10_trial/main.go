package main

import "fmt"

func main() {

	x := 40

	if x > 50 {
		fmt.Println("x is greater than 50.")
	} else if x == 50 {
		fmt.Println("x is equal to 50.")
	} else {
		fmt.Println("x is lower than 50")
	}
	fmt.Printf("The actual value  of x: %v\n", x)
}
