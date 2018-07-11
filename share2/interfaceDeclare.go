package main

import "fmt"

//startOMIT
type Stringer interface {
	String() string
}
type Printer interface {
	Stringer
	Print()
}
type User struct {
	id   int
	name string
}

func (u *User) String() string {
	return fmt.Sprintf("User %v, %v", u.name, u.id)
}
func (u *User) Print() {
	fmt.Println(u.String())
}
func main() {
	var u Printer = &User{id: 1, name: "zhengkeda"}
	u.Print()
	//endOMIT
}
