package domain

import (
	"context"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func GetTracerSpan(ctx context.Context, adapter string, spanBase string, name string) (trace.Tracer, context.Context, trace.Span) {
	tr := global.Tracer(adapter)
	newCtx, span := tr.Start(ctx, spanBase+name)

	return tr, newCtx, span
}
