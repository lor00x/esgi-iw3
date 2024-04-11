package main

import (
	"fmt"
)

type BoxString struct {
	content string
}

type BoxUser struct {
	content User
}

type Box[T any] struct {
	content T
}

func (b Box[T]) Content() T {
	return b.content
}

type User struct {
	Name string
}

func main() {
	b := Box[string]{content: "Hello"}
	b2 := Box[User]{content: User{Name: "John"}}
	fmt.Println("Content:", b.content)
	fmt.Println("Content 2:", b2.content)
}
