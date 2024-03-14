package main

import "fmt"

func main() {
	var a = 42
	b := &a

	// a = 99
	*b = 99

	fmt.Printf("%#v", a)

}
