package main

import "fmt"

func Print(i interface{}) {
	fmt.Printf("%v", i)
}

func main() {
	var i int = 11
	Print(i)
	var s string = "str"
	Print(s)
}
