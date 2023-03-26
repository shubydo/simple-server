package server

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

func (s *server) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info(
			"Hello page hit",
			zap.String("method", r.Method),
			zap.String("url", r.URL.String()),
		)

		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}

		fmt.Fprintf(w, "Hello!")

		return
	}
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info(
			"Index page hit",
			zap.String("method", r.Method),
			zap.String("url", r.URL.String()),
		)

		w.Write([]byte("Index!"))
	}
}
