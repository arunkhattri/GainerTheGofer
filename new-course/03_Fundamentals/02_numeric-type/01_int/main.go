package main

import "fmt"

var x = 42
var a int8 = -128 // the set of all signed 8-bit integers (-128 to 127)

func main() {
	y := 11214
	z := 5.28
	fmt.Printf("%d, %T\n", x, x)
	fmt.Printf("%d, %T\n", y, y)
	fmt.Printf("%f, %T\n", z, z)
	fmt.Printf("%d, %T\n", a, a)
}
