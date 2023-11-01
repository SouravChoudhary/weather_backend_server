package middleware

import (
	"net/http"
	"time"
	"weatherapp/pkg/logging"

	"github.com/google/uuid"
)

// LoggingMiddleware is a middleware function that logs incoming HTTP requests.
func LoggingMiddleware(logger logging.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			requestID := uuid.New()
			next.ServeHTTP(w, r)
			duration := time.Since(startTime).Milliseconds()
			logger.LogInfo("Incoming request",
				map[string]interface{}{
					"requestID":  requestID.String(),
					"method":     r.Method,
					"path":       r.URL.Path,
					"remoteAddr": r.RemoteAddr,
					"duration":   duration,
				},
			)
		})
	}
}
