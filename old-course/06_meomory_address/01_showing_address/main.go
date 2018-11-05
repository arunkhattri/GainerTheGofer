package main

import "fmt"

func main() {
	a := 43

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a) // & for memory address
	fmt.Printf("%d\n", &a)
}
