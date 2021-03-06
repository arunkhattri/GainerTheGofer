package main
import "fmt"

// one way

// var name = "Arun Kr. Khattri"
//
// func main() {
// 	fmt.Println("Hello", name)
// }

// another way
// func main() {
// 	var name = "Arun Kr. Khattri"
// 	fmt.Println("Hello", name)
// }

// best way

// func main() {
// 	name := "Arun Kr. Khattri"
// 	fmt.Println("Hello", name)
// }

// another way

func main() {
	name := `Arun Kr. Khattri` // back-ticks work like double-quotes
	fmt.Println("Hello", name)
}
