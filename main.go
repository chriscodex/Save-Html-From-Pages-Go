package main

import (
	"fmt"
	"net/http"
)

func main() {
	response, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(response)
	}
}
