package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Pass to the next handler
		next.ServeHTTP(w, r)

		// After handler runs
		log.Printf("[%s] %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}