package main

import "fmt"

func main() {

	iterateur := NewIterateur()
	one := iterateur()
	two := iterateur()
	three := iterateur()
	fmt.Println(one, two, three, iterateur(), iterateur())
}

func NewIterateur() func() int {
	i := 0 // Etat interne
	return func() int {
		i++
		return i
	}
}
