package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	age  int    `json:"age"`
}

func main() {
	marshal()
}
func marshal() {
	var stu = &Student{Name: "zhengkeda", age: 18}
	b, err := json.Marshal(stu)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(b))
	var stu2 = &Student{}
	err = json.Unmarshal(b, stu2)
	fmt.Println(*stu2)
}
