package main

import (
	"fmt"
	"net/http"
)

func main() {
	if err := http.ListenAndServe("127.0.0.1:9090", http.FileServer(http.Dir("../../assets"))); err != nil {
		fmt.Println("server start failed")
		return
	}
}
