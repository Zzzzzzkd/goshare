package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	m := map[string]string{}
	go func() {
		for i := 0; i < 200000; i++ {
			istr := strconv.Itoa(i)
			m[istr] = istr
		}
	}()

	go func() {
		for i := 0; i < 200000; i++ {
			if _, ok := m["02"]; ok {
				fmt.Println("aa")
			}
		}
	}()
	time.Sleep(3 * time.Second)
}
