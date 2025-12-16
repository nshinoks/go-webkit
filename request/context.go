package request

import "context"

type ctxKey int

const requestIDKey ctxKey = iota

func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey, id)
}

func RequestIDFrom(ctx context.Context) (string, bool) {
	v := ctx.Value(requestIDKey)
	s, ok := v.(string)
	return s, ok
}
