package main

import "fmt"

func main() {
	a := [2]int{}
	fmt.Println(a)

	s1 := []int{2, 4, 6}
	s2 := []int{3, 5, 7}
	s1 = append(s1, 1)
	fmt.Println(s1)
	s1 = append(s1, s2...)
	fmt.Println(s1)
	fmt.Println(s1[1:3])
}
