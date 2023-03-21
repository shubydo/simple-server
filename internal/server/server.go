package server

import "net/http"

type server struct {
	router *http.ServeMux
	logger interface{}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func newServer() *server {
	s := &server{
		router: http.NewServeMux(),
	}
	s.routes()

	return s
}

func NewServer() *server {
	s := newServer()
	return s
}
