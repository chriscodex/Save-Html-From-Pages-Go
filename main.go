package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type writerWeb struct{}

func (writerWeb) Write(p []byte) (int, error) {
	ioutil.WriteFile("index.html", p, 0644)
	return len(p), nil
}

func main() {
	response, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
	}
	e := writerWeb{}
	io.Copy(e, response.Body)
}
