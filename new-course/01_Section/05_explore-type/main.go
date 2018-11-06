package main

import "fmt"

var x = 11214
var y = "Shaken, not stirred"                    // Interpreted string literal
var z = `James Bond said, "Shaken, not stirred"` // Raw string literals

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
