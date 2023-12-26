package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Create a response writer that captures the status code
		rw := &responseWriter{ResponseWriter: w}

		// Call the next handler in the chain
		next.ServeHTTP(rw, r)

		// Log information about the request including the status code
		duration := time.Since(startTime)
		log.Printf("[%s] %s %s %d %s", r.Method, r.RequestURI, r.RemoteAddr, rw.status, duration)
	})
}

// Custom response writer to capture the status code
type responseWriter struct {
	http.ResponseWriter
	status int
}

// Implement WriteHeader to capture the status code
func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}
