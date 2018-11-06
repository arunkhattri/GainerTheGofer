package main

import "fmt"

func main() {
	fmt.Println("Hello everyone, This is Golang Programming language")

	foo()

	fmt.Println("something more")

	for i := 0; i < 100; i++ {
		if i%25 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("I'am in a foo.")
}

func bar() {
	fmt.Println("and then we exited")
}
