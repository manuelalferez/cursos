package main

import (
	"net/http"
)

/**
	Maneja las URLs que nos llegan
 */
type Router struct {
	rules map[string]http.HandlerFunc // cada string (url) le asociamos un Handler
}

func newRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) FindHandler(path string) (http.HandlerFunc, bool){
	handler, error := r.rules[path]
	return handler, error
}

func (r *Router) ServeHTTP(w http.ResponseWriter, responde *http.Request) {
	handler, exist := r.FindHandler(responde.URL.Path)
	if !exist{
		w.WriteHeader(http.StatusNotFound)
	}else{
		handler(w, responde)
	}
}
