package server

import "net/http"

type server struct {
	router *router
	// db *someDB
	// logger *logger
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func newServer() *server {
	s := &server{}
	s.routes()

	return s
}

// package server

// // func initHandlers() {
// // 	http.HandleFunc("/hello", handleHello)
// // }

// type ServerOption func(*Server)

// type Server struct {
// 	// Addr   string
// 	Port int
// 	// Router *Router
// 	// Logger *Logger
// }

// func NewServer(options ...ServerOption) *Server {
// 	server := &Server{
// 		// Addr:   ":8080",
// 		Port: 8080,
// 		// Router: NewRouter(),
// 		// Logger: NewLogger(),
// 	}

// 	for _, option := range options {
// 		option(server)
// 	}

// 	return server
// }
