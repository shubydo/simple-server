package server

import "github.com/prometheus/client_golang/prometheus/promhttp"

func (s *Server) routes() {
	// Main routes
	s.router.HandleFunc("/", s.handleIndex())
	s.router.HandleFunc("/hello", s.handleHello())
	s.router.HandleFunc("/healthz", s.handleHealthz())
	s.router.HandleFunc("/readyz", s.handleReadyz())

	// Prometheus metrics
	s.router.Handle("/metrics", promhttp.Handler())
}
