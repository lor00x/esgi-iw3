package main

import (
	"fmt"
	"strings"
)

type Strategy func(string)

var uppercase = func(msg string) {
	fmt.Println(strings.ToUpper(msg))
}

var lowercase = func(msg string) {
	fmt.Println(strings.ToLower(msg))
}

func main() {

	msg := "Salut ESGI 3!"
	var s Strategy
	s = lowercase
	Display(msg, s)
}

func Display(msg string, strat Strategy) {
	strat(msg)
}
