package main

import "fmt"

type User struct {
	name string
	Age  int
}

func (u *User) SetName(n string) {
	u.name = n
}

// Oops, forgot the ‘*’
func (u *User) SetAge(a int) {
	u.Age = a
}

func main() {
	u := &User{}

	u.SetAge(10)
	//u.name = "Harry"
	//u.SetName("Harry") // OK
	fmt.Printf("%#v", u)
	// u.SetAge(16)       // Useless

}
