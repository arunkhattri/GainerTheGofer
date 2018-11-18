package main

import "fmt"

func main() {
	x := 5
	y := 14
	g := y == x
	h := y <= x
	i := y >= x
	j := y != x
	k := y < x
	l := y > x

	fmt.Printf("g is %t\n", g)
	fmt.Printf("h is %t\n", h)
	fmt.Printf("i is %t\n", i)
	fmt.Printf("j is %t\n", j)
	fmt.Printf("k is %t\n", k)
	fmt.Printf("l is %t\n", l)

}
