package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//ifelsebeginOMIT
	var b bool
	if b {
		fmt.Println("b is true")
	} else {
		fmt.Println("b is false")
	}
	m := map[string]string{"name": "zhengkeda"}
	if b, err := json.Marshal(m); err != nil {
		fmt.Println("json error:", err)
	} else {
		fmt.Println(string(b))
	}
	//ifelseendOMIT

	//forbeginOMIT
	for i := 0; i < 5; i++ {
		if i < 2 {
			continue
		}
		if i > 3 {
			break
		}
		fmt.Println(i)
	}
	//forendOMIT

	//forrangebeginOMIT
	s := []int{1, 2, 3, 4, 5, 6, 7}
	for k, v := range s {
		fmt.Printf("k=%v, v=%v\n", k, v)
	}
	//forrangeendOMIT

	//switchbeginOMIT
	var a = "a"
	switch a {
	case "a":
		fmt.Println("aaaa")
	case "b":
		fmt.Println("bbbb")
	}
	//switchendOMIT

	//gotobeginOMIT
	var j = 0
loop:
	j++
	if j == 5 {
		return
	}
	goto loop
	//gotoendOMIT

}
