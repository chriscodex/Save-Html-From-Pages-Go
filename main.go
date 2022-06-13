package main

import (
	"fmt"
	"io"
	"net/http"
)

type writerWeb struct{}

func (writerWeb) Write(p []byte) (int, error) {
	fmt.Println(p)
	return 0, nil
}

func main() {
	response, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
	}
	e := writerWeb{}
	io.Copy(e, response.Body)
}
