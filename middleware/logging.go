package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)

		slog.Info(
			"Health Check",
			slog.Int("Method", wrapped.statusCode),
			r.Method,
			r.URL.Path,
			slog.String("time", time.Since(start).String()),
		)
		// log.Println("INFO", wrapped.statusCode, r.Method, r.URL.Path, time.Since(start))
	})
}
