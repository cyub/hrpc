package ctxsetters

import (
	"context"
	"net/http"
	"strconv"

	"github.com/cyub/hrpc/internal/contextkeys"
)

func WithMethodName(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, contextkeys.MethodNameKey, name)
}

func WithServiceName(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, contextkeys.ServiceNameKey, name)
}

func WithPackageName(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, contextkeys.PackageNameKey, name)
}

func WithStatusCode(ctx context.Context, code int) context.Context {
	return context.WithValue(ctx, contextkeys.StatusCodeKey, strconv.Itoa(code))
}

func WithResponseWriter(ctx context.Context, w http.ResponseWriter) context.Context {
	return context.WithValue(ctx, contextkeys.ResponseWriterKey, w)
}
