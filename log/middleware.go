package log

import (
	"context"

	"github.com/rs/xid"
)

type contextKey string

const contextKeyRequestID = contextKey("log.req-id")

func InitLoggingContext(ctx context.Context) context.Context {
	ctx = setRequestID(ctx)
	return ctx
}

func setRequestID(ctx context.Context) context.Context {
	if GetRequestID(ctx) != "" {
		return ctx
	}

	ctx = context.WithValue(ctx, contextKeyRequestID, xid.New().String())
	return ctx
}

func GetRequestID(ctx context.Context) string {
	reqID, ok := ctx.Value(contextKeyRequestID).(string)
	if !ok {
		return ""
	}

	return reqID
}
