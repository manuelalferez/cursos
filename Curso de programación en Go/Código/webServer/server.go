package main

import "net/http"

/**
	Crea una conexión y escucha
 */
type Server struct {
	port   string
	router *Router
}

func newServer(port string) *Server {
	return &Server{
		port:   port,
		router: newRouter(),
	}
}

func (s *Server) listen() error {
	http.Handle("/", s.router)
	error := http.ListenAndServe(s.port, nil)
	if error != nil {
		return error
	}
	return nil
}