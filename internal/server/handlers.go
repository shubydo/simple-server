package server

import (
	"fmt"
	"net/http"
)

func (s *server) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("Hello page hit")
		if r.Method == "GET" {
			fmt.Fprintf(w, "Hello!")
			return
		}

		// Return "Method not allowed" if not a GET request

		fmt.Fprintf(w, "Method not allowed")

		return
	}
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("Index page hit")
		w.Write([]byte("Index!"))
	}
}
