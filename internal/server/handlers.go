package server

import (
	"net/http"

	"go.uber.org/zap"
)

func (s *Server) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// s.logger.Info(
		// 	"handleHello called",
		// 	zap.String("method", r.Method),
		// 	zap.String("url", r.URL.String()),
		// 	zap.Any("headers", r.Header),
		// )

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			_, err := w.Write([]byte(`Method not allowed`))
			if err != nil {
				s.logger.Error("Error writing response", zap.Error(err))
				return
			}
			return
		}

		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`Hello!`))
		if err != nil {
			s.logger.Error("Error writing response", zap.Error(err))
		}
	}
}

func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info(
			"handleIndex called",
			zap.String("method", r.Method),
			zap.String("url", r.URL.String()),
			zap.Any("headers", r.Header),
		)

		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Index!"))
		if err != nil {
			s.logger.Error("Error writing response", zap.Error(err))
			return
		}
	}
}

//nolint:unused
func (s *Server) handleNotFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info(
			"handleNotFound called",
			zap.String("method", r.Method),
			zap.String("url", r.URL.String()),
		)

		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte("404!"))
		if err != nil {
			s.logger.Error("Error writing response", zap.Error(err))
			return
		}
	}
}

func (s *Server) handleHealthz() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info(
			"handleHealthz called",
			zap.String("method", r.Method),
			zap.String("url", r.URL.String()),
		)

		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("OK"))
		if err != nil {
			s.logger.Error("Error writing response", zap.Error(err))
			return
		}
	}
}

func (s *Server) handleReadyz() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		_, err := w.Write([]byte("OK"))
		if err != nil {
			s.logger.Error("Error writing response", zap.Error(err))
			return
		}
	}
}

// func (s *server) metricsHandler() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// w.WriteHeader(http.StatusOK)
// 		// w.Write([]byte("Metrics!"))
// 	}
// }

// logMiddleware is a middleware that logs the request.
func (s *Server) logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info(
			"Request received",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
		)

		s.router.ServeHTTP(w, r)
	})

	// return func(next http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		s.logger.Info(
	// 			"Request received",
	// 			zap.String("method", r.Method),
	// 			zap.String("path", r.URL.Path),
	// 		)

}
