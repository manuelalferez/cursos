package main

import (
	"net/http"
)

/**
	Maneja las URLs que nos llegan
 */
type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func newRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool){
	_, pathExist := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return handler, methodExist, pathExist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, response *http.Request) {
	handler, methodExist, pathExist := r.FindHandler(response.URL.Path, response.Method)
	if !pathExist{
		w.WriteHeader(http.StatusNotFound)
	}else if !methodExist{
		w.WriteHeader(http.StatusMethodNotAllowed)
	}else{
		handler(w, response)
	}
}
