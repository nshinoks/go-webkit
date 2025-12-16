package middleware

import (
	"net/http"
	"time"
)

type AccessLogFunc func(r *http.Request, status int, dur time.Duration)

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func AccessLog(fn AccessLogFunc) func(http.Handler) http.Handler {
	if fn == nil {
		fn = func(_ *http.Request, _ int, _ time.Duration) {}
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sw := &statusWriter{ResponseWriter: w, status: http.StatusOK}
			start := time.Now()
			next.ServeHTTP(sw, r)
			fn(r, sw.status, time.Since(start))
		})
	}
}
