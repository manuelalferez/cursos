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

func (s *Server) setHandler(method string, path string, handle http.HandlerFunc){
	_, pathExist:= s.router.rules[path]
	if !pathExist{
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handle
}

func (s *Server) addMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc{
	for _, middleware := range middlewares{
		f = middleware(f)
	}
	return f
}

func (s *Server) listen() error {
	http.Handle("/", s.router)
	error := http.ListenAndServe(s.port, nil)
	if error != nil {
		return error
	}
	return nil
}