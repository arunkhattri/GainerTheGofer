package main

import "fmt"

const (
	a = iota //0
	// b = iota //1
	// c = iota //2
	// even this will do
	b
	c
)

// will reset the iota
const (
	d = iota //0
	e        //1
	f        //2
)
const (
	_ = iota      //0
	B = iota * 10 // 1 * 10
	C = iota * 10 // 2 * 10
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(B)
	fmt.Println(C)
}
