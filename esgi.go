package main

import "fmt"

func main() {
	var a int = 42

	var b *int = &a

	var c = &b // **int

	fmt.Println(b)

	// Déréférencement
	fmt.Println(*b)
}
