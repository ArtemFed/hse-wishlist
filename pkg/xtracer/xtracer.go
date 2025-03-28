package xtracer

import (
	"context"
	"fmt"
	"github.com/ArtemFed/hse-wishlist/pkg/xapp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func Init(cfg *Config, appCfg *xapp.Config) (*sdktrace.TracerProvider, error) {
	ctx := context.Background()
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(appCfg.Name),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	var batcher sdktrace.TracerProviderOption
	if cfg.Enable {
		log.Print("[Tracer] Init started as Enabled")
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		log.Print("cfg.ExpTarget", cfg.ExpTarget)
		conn, err := grpc.DialContext(ctx,
			cfg.ExpTarget,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock(),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
		}

		exp, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
		if err != nil {
			return nil, fmt.Errorf("failed to create trace exporter: %w", err)
		}
		batcher = sdktrace.WithBatcher(exp)
	} else {
		log.Print("[Tracer] Init started as Disabled")
		exp, err := stdout.New(stdout.WithPrettyPrint())
		if err != nil {
			return nil, fmt.Errorf("failed to create trace exporter: %w", err)
		}
		batcher = sdktrace.WithBatcher(exp)
	}

	log.Print("[Tracer] Creating NewTracerProvider")
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		batcher,
		sdktrace.WithResource(res),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}
