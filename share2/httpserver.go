package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", f)
	if err := http.ListenAndServe(":1001", nil); err != nil {
		fmt.Println(err.Error)
	}
}

func f(w http.ResponseWriter, r *http.Request) {
}
