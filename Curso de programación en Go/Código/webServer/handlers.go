package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func HandlerAPI(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "I am an API")
}

func PostRequest(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var metaData MetaData
	error := decoder.Decode(&metaData)
	if error != nil {
		fmt.Fprintf(w, "error: %v", error)
		return
	}else{
		fmt.Fprintf(w, "Payload: %v", metaData)
	}
}

func UserPostRequest(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var user User
	error := decoder.Decode(&user)
	if error != nil {
		fmt.Fprintf(w, "error: %v", error)
		return
	}else{
		response, error := user.toJSON()
		if error != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}