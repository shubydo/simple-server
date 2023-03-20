package server

import (
	"fmt"
	"net/http"
)

func (s *server) handleHello(w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			fmt.Fprintf(w, "Hello!")
			return
		}

		// Return "Method not allowed" if not a GET request
		fmt.Fprintf(w, "Method not allowed")
	}
}
