package server

import "github.com/prometheus/client_golang/prometheus/promhttp"

func (s *Server) routes() {
	// Prometheus metrics
	s.router.Handle("/metrics", promhttp.Handler())

	// Main routes
	s.router.HandleFunc("/", s.handleIndex())
	s.router.HandleFunc("/hello", s.handleHello())
	s.router.HandleFunc("/healthz", s.handleHealthz())
}
