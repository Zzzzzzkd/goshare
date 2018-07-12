package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	//osRead()
	osWrite()
	ioutilReadWrite()
}

func osRead() {
	f, err := os.Open("/etc/redis/6379.conf")
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()
	b := make([]byte, 1024)
	for {
		_, err := f.Read(b)
		fmt.Printf("%v", string(b))
		if err != nil {
			if err != io.EOF {
				panic(err.Error())
			} else {
				return
			}
		}
	}
}

func osWrite() {
	f, err := os.Create("/tmp/osWrite")
	if err != nil {
		panic(err.Error())
	}
	_, err = f.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))
	if err != nil {
		panic(err.Error())
	}
}

func ioutilReadWrite() {
	b, err := ioutil.ReadFile("/etc/redis/6379.conf")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(len(b))

	err = ioutil.WriteFile("/tmp/ioutiltest", []byte(time.Now().Format("2006-01-02 15:04:05")), 0777)
	if err != nil {
		panic(err.Error())
	}
}
