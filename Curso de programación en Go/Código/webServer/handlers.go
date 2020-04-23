package main

import (
	"fmt"
	"net/http"
)

func HandlerRoot(w http.ResponseWriter, response *http.Response){
	fmt.Fprintln(w, "Hello World!")
}