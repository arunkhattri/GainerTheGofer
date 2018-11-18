package main

import (
	"fmt"
)

func main() {
	name := "Aaryaansh Arora Khattri"
	fmt.Printf("%s, %T\n", name, name)

	bs := []byte(name)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(name); i++ {
		fmt.Printf("%#U ", name[i])

		fmt.Println("")

		for i, v := range name {
			fmt.Printf("Index: %d, Hex: %X\n", i, v)
		}
	}
}
