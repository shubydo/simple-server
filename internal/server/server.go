package server

import (
	"net/http"

	"go.uber.org/zap"
)

const (
	// DefaultPort is the default port for the server.
	DefaultPort = 7777
)

type Server struct {
	router *http.ServeMux
	logger *zap.Logger
	port   int
}

// Option is a function that configures the server.
type Option func(*Server)

// WithLogger sets the logger for the server.
//
//goland:noinspection GoUnusedExportedFunction
func WithLogger(logger *zap.Logger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}

// WithPort sets the port for the server.
func WithPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

// New returns a new instance of the server.
func New(opts ...Option) *Server {
	s := &Server{
		router: http.NewServeMux(),
		logger: newLogger(),
		port:   DefaultPort,
	}

	for _, opt := range opts {
		opt(s)
	}

	s.routes()

	return s
}

func newLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer func() {
		_ = logger.Sync()
	}()

	return logger
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Request received", zap.String("method", r.Method), zap.String("path", r.URL.Path))
	s.router.ServeHTTP(w, r)
}
