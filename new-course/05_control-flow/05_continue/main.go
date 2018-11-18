// to show how keyword continue work
// print even numbers
package main

import "fmt"

func main() {
	x := 0
	// initialize for loop
	for {
		x++

		// condition for break
		if x > 10 {
			break
		}
		if x%2 != 0 {
			// condition for continue
			continue
		}
		fmt.Println(x)
	}
}
