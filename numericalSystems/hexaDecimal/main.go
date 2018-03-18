package main

import "fmt"

func main() {
	// add leading 0 for hex (0x)
	fmt.Printf("%d in Hexadecimal: %#x\n", 42, 42)
	// add leading 0 for hex (0X)
	fmt.Printf("%d in Hexadecimal: %#X\n", 42, 42)
	// without leading 0
	fmt.Printf("%d in Hexadecimal: %x\n", 42, 42)
}
