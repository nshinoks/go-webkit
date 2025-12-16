package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"
)

type ctxKey int

const requestIDKey ctxKey = iota

type RequestIDOptions struct {
	HeaderName string
	Generator  func() string
}

type RequestIDOption func(*RequestIDOptions)

func WithHeader(name string) RequestIDOption {
	return func(o *RequestIDOptions) { o.HeaderName = name }
}
func WithGenerator(gen func() string) RequestIDOption {
	return func(o *RequestIDOptions) { o.Generator = gen }
}

func RequestID(opts ...RequestIDOption) func(http.Handler) http.Handler {
	o := &RequestIDOptions{
		HeaderName: "X-Request-Id",
		Generator: func() string {
			var b [16]byte
			_, _ = rand.Read(b[:])
			return hex.EncodeToString(b[:])
		},
	}
	for _, opt := range opts {
		opt(o)
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id := r.Header.Get(o.HeaderName)
			if id == "" {
				id = o.Generator()
			}
			w.Header().Set(o.HeaderName, id)
			ctx := context.WithValue(r.Context(), requestIDKey, id)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RequestIDFrom(ctx context.Context) (string, bool) {
	v := ctx.Value(requestIDKey)
	s, ok := v.(string)
	return s, ok
}
