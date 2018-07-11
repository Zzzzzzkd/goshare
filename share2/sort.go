package main

import (
	"fmt"
	"sort"
)

//startOMIT
type Student struct {
	Id   int
	Name string
}
type StudentList []Student

func (sl StudentList) Swap(i, j int) {
	sl[i], sl[j] = sl[j], sl[i]
}
func (sl StudentList) Len() int {
	return len(sl)
}
func (sl StudentList) Less(i, j int) bool {
	return sl[i].Id < sl[j].Id
}

func main() {
	var sl StudentList = StudentList{}
	sl = append(sl, Student{2, "zhangsan"})
	sl = append(sl, Student{3, "lisi"})
	sl = append(sl, Student{1, "wangwu"})
	fmt.Printf("%+v\n", sl)
	sort.Sort(sl)
	fmt.Printf("%+v\n", sl)
}

//endOMIT
