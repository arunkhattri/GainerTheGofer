package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// reading an integer
	var age int
	fmt.Println("What is your age?")
	_ , err:= fmt.Scan(&age)
	\\ use err

	// reading a string
	var name string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	name, _ = reader.ReadString('\n')

	fmt.Println("Your name is ", name, " and your age is ", age)
}
