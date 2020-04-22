package main

import (
	"fmt"
	"io"
	"net/http"
)

type webWriter struct{}

func (webWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	response, error := http.Get("http://google.com")
	webW := webWriter{}
	if error != nil {
		fmt.Println(error)
	}
	io.Copy(webW, response.Body)
}
