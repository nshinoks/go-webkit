package respond

import (
	"net/http"
	"time"

	"github.com/nshinoks/go-webkit/request"
)

type Meta struct {
	RequestID string    `json:"requestId,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

type Envelope[T any] struct {
	Data T    `json:"data"`
	Meta Meta `json:"meta,omitempty"`
}

type Options struct {
	// Whether to include timestamp in meta. Default: true
	IncludeTimestamp bool
	// Whether to include requestId in meta (from context). Default: true
	IncludeRequestID bool
	// Clock for timestamp (testability). Default: time.Now().UTC()
	Now func() time.Time
}

type Option func(*Options)

func WithTimestamp(enabled bool) Option {
	return func(o *Options) { o.IncludeTimestamp = enabled }
}
func WithRequestID(enabled bool) Option {
	return func(o *Options) { o.IncludeRequestID = enabled }
}
func WithNow(now func() time.Time) Option {
	return func(o *Options) { o.Now = now }
}

func defaultOptions() Options {
	return Options{
		IncludeTimestamp: true,
		IncludeRequestID: true,
		Now:              func() time.Time { return time.Now().UTC() },
	}
}

// OK writes a 200 response with unified envelope.
func OK[T any](w http.ResponseWriter, r *http.Request, data T, opts ...Option) {
	Write(w, r, http.StatusOK, data, opts...)
}

// Created writes a 201 response with unified envelope.
func Created[T any](w http.ResponseWriter, r *http.Request, data T, opts ...Option) {
	Write(w, r, http.StatusCreated, data, opts...)
}

// NoContent writes a 204 response.
func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func Write[T any](w http.ResponseWriter, r *http.Request, status int, data T, opts ...Option) {
	o := defaultOptions()
	for _, opt := range opts {
		opt(&o)
	}

	env := Envelope[T]{Data: data}
	if r != nil {
		if o.IncludeRequestID {
			if rid, ok := request.RequestIDFrom(r.Context()); ok {
				env.Meta.RequestID = rid
			}
		}
	}
	if o.IncludeTimestamp {
		env.Meta.Timestamp = o.Now()
	}

	JSON(w, status, env)
}
