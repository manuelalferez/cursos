package main

import (
	"fmt"
	"net/http"
)

func HandlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func HandlerAPI(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "I am an API")
}