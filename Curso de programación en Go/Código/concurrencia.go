package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkServer(server string, channel chan string) {
	_, error := http.Get(server)
	if error != nil {
		fmt.Printf(server, "is not working :(\n")
		channel <- server + "is not working :(\n"
	} else {
		fmt.Printf(server, "is working :)\n")
		channel <- server + "is working :)\n"
	}
}

func main() {
	startTime := time.Now()
	channel := make(chan string)
	servers := []string{
		"http://google.com",
		"http://twitter.com",
		"http://github.com",
		"http://youtube.com",
	}
	counter := 0
	for {
		if counter > 5 {
			break
		}
		for _, server := range servers {
			go checkServer(server, channel)
		}
		time.Sleep(5 * time.Second)
		counter++
	}

	for i := 0; i < len(servers); i++ {
		fmt.Println(<-channel)
	}
	diffTime := time.Since(startTime)
	fmt.Println(diffTime)
}
