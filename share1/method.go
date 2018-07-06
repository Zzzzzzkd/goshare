package main

import "fmt"

//declarebeginOMIT
type S struct {
}

func (s *S) Println() {
	fmt.Println("this is s print")
}

func (s S) Println2() {
	fmt.Println("this is s print2")
}

//declareendOMIT
//extendbeginOMIT
type Mouse struct {
}

func (m Mouse) eat() {
	fmt.Println("eating eating")
}

type Animal struct {
	Mouse
}

func main() {
	a := Animal{}
	a.eat()
}

//extendendOMIT
