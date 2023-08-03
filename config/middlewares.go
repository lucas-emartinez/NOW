package config

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(handler http.HandlerFunc) http.HandlerFunc

// LoggingMiddleware logs the incoming HTTP request.
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Start timer
		startTime := time.Now()

		// Process request
		next(w, r)

		// Stop timer
		elapsedTime := time.Since(startTime)

		// Log request details
		const (
			colorReset  = "\033[0m"
			colorRed    = "\033[31m"
			colorGreen  = "\033[32m"
			colorYellow = "\033[33m"
		)
		log.Printf(
			"[%s%s%s] %s%s%s %s%s%s %s",
			colorYellow, r.Method, colorReset,
			colorGreen, r.RequestURI, colorReset,
			colorRed, r.RemoteAddr, colorReset,
			elapsedTime,
		)
	}
}
