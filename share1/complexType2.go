package main

import "fmt"

func main() {
	m := map[string]string{
		"name": "zhengkeda",
		"age":  "18",
	}
	fmt.Println(m["age"])
	delete(m, "age")
	fmt.Println(m)

	type Student struct {
		Age  int
		Name string
	}
	zhengkeda := &Student{Age: 0, Name: "zhengkeda"}
	fmt.Println(zhengkeda.Age)
}
