package server

import (
	"net/http"

	"go.uber.org/zap"
)

type server struct {
	router *http.ServeMux
	logger *zap.Logger
}

func newLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer func() {
		_ = logger.Sync()
	}()

	return logger
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func newServer() *server {
	s := &server{
		router: http.NewServeMux(),
		logger: newLogger(),
	}
	s.routes()

	return s
}

func NewServer() *server {
	s := newServer()
	return s
}
