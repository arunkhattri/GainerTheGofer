package main

import "fmt"

func main() {
	x := 14
	// print x in decimal, binary and hex
	fmt.Printf("%d\t\t%b\t\t%#x\n", x, x, x)

	// shifts the bits of that int over 1 position to the left,
	// and assigns that to a variable
	y := x << 1
	fmt.Printf("%d\t\t%b\t\t%#x\n", y, y, y)
}
