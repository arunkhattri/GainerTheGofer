package main

import "fmt"

const (
	py  = 2018 + iota
	ny  = 2018 + iota
	ny2 = 2018 + iota
	ny3 = 2018 + iota
)

func main() {
	fmt.Println(py, ny, ny2, ny3)
}
