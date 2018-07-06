package main

import "fmt"

//int float64 bool string; struct array slice map channel

var s string
var i int = 0
var j = 0

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println(s, i, j)
	var I int = 0
	fmt.Println(I)

	x := "this is string"
	x = `is string`
	fmt.Println(x)

	xiaoming := Student{
		Name: "xiaoming",
		Age:  18,
	}
	fmt.Println(xiaoming)

	//var array [4]int
	//array = [4]int{1, 2, 3, 4}
	//array := new([4]int)
	array := [4]int{}
	fmt.Println(array)

	//slice := []string{"a", "bc"}
	//slice := make([]string, 0)
	var slice []string
	slice = []string{"a", "b"}
	fmt.Println("slice is ", slice)

	//var m map[string]Student
	//m = map[string]Student{
	//	"a": xiaoming,
	//}
	//m := make(map[string]Student, 0)
	m := map[string]Student{
		"a": xiaoming,
	}
	fmt.Println("m is ", m)

	var c chan int
	c = make(chan int)
	fmt.Println(c)
}
