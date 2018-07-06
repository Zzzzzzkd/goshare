package main

import "fmt"

func funcName() {
}

func mutRet() (string, int) {
	return "", 0
}

func retPreDeclare() (x string, y int) {
	return
	//return "",1
}

func changeArg(ss ...int) {
}

func deferfunc() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
	return
}

func panicfunc() {
	defer func() {
		if perr := recover(); perr != nil {
			fmt.Println(perr)
		}
	}()
	panic("ppppppp")
}
func main() {
	deferfunc()
	panicfunc()
}
