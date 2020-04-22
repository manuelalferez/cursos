package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkServer(server string) {
	_, error := http.Get(server)
	if error != nil {
		fmt.Printf(server, "is not working :(\n")
	} else {
		fmt.Printf(server, "is working :)\n")
	}
}

func main() {
	startTime := time.Now()
	servers := []string{
		"http://google.com",
		"http://twitter.com",
		"http://github.com",
		"http://youtube.com",
	}
	for _, server := range servers {
		checkServer(server)
	}
	diffTime := time.Since(startTime)
	fmt.Println(diffTime)
}
