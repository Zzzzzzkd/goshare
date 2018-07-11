package main

import "os"

func main() {
}

func osfile() {
	f, err := os.Open("/etc/redis/6379.conf")
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()
	b := make([]byte, 0, 1024)
	for {
		n, err := f.Read(b)
		if err 
	}
}
