package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("../../assets")))
	if err := http.ListenAndServe("127.0.0.1:9090", nil); err != nil {
		fmt.Println("server start failed")
		return
	}
}
