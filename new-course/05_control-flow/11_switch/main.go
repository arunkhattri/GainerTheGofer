package main

import (
	"fmt"
	"time"
)

func main() {
	// Basic switch with default
	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("Today is Saturday")
	case time.Sunday:
		fmt.Println("Today is Sunday")
	default:
		fmt.Println("Today is weekday")
	}
	fmt.Printf("Today is: %v\n", time.Now().Weekday())
}
