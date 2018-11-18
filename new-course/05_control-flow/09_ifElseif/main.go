package main

import "fmt"

func main() {
	x := 14
	if x == 20 {
		fmt.Println("Our value was 20")
	} else if x == 15 {
		fmt.Println("Our value was 15")
	} else {
		fmt.Printf("Our value was %d", x)
	}
}
