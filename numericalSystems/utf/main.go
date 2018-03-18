package main

import "fmt"

func main() {
	for i := 65; i < 91; i++ {
		fmt.Printf("Decimal: %d \t Binary: %b \t UTF: %q \n", i, i, i)
	}
}
