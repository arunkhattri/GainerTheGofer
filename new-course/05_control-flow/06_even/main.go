// Print even numbers
package main

import "fmt"

func main() {
	x := 1

	for {
		if x > 10 {
			break
		}
		if x%2 == 0 {
			fmt.Printf("%d : Even Number\n", x)
		}
		x++
	}
}
