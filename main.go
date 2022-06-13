package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type writerWeb struct{}

func (writerWeb) Write(p []byte) (int, error) {
	ioutil.WriteFile("index.html", p, 0644)
	return len(p), nil
}

func scan() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	url := scan()
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	e := writerWeb{}
	io.Copy(e, response.Body)
}
