package main

import "fmt"

// Type nommé
type Operation func(int, int) int

// Type func(int, int) int
func Add(a, b int) int {
	return a + b
}
func Sub(x, z int) int {
	return x - z
}

// Operation et func(int, int) int
// sont compatibles
var defaultOp Operation = Add

func main() {
	// Appel de la fonction assignée
	// à la variable defaultOp
	defaultOp = Sub
	res := defaultOp(42, 90)
	fmt.Println(RunOperation(43, 45, Sub))
	fmt.Printf("%#v", res)
}

func RunOperation(a, b int, op Operation) int {
	return op(a, b)
}
