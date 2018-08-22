package main

import (
	"context"
	"github.com/batazor/go-logger/pkg/amqp"
	"github.com/batazor/go-logger/pkg/grpc"
	"github.com/batazor/go-logger/pkg/healthcheck"
	"github.com/batazor/go-logger/pkg/jaeger"
	"github.com/batazor/go-logger/pkg/metrics"
	"github.com/batazor/go-logger/pkg/redis"
	"github.com/batazor/go-logger/utils"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

var (
	// Logging
	log = logrus.New()

	// ENV
	REDIS_ENABLE        = utils.Getenv("REDIS_ENABLE", "true")
	AMQP_ENABLE         = utils.Getenv("AMQP_ENABLE", "true")
	PROMETHEUS_ENABLE   = utils.Getenv("PROMETHEUS_ENABLE", "true")
	GRPC_ENABLE         = utils.Getenv("GRPC_ENABLE", "true")
	OPENTRACING_ENABLED = utils.Getenv("OPENTRACING_ENABLED", "true")
)

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)

	// Global context
	ctx := context.Background()

	// OpenTracing =============================================================
	if OPENTRACING_ENABLED == "true" {
		tracer, closer := jaeger.Listen()
		opentracing.SetGlobalTracer(tracer)
		defer closer.Close()

		// Add event
		span := tracer.StartSpan("say-hello")
		ctx = opentracing.ContextWithSpan(ctx, span)
		jaeger.Add(ctx)
	}
}

func main() {
	// Run REDIS
	if REDIS_ENABLE == "true" {
		go redis.Listen()
	}

	// Run AMQP
	if AMQP_ENABLE == "true" {
		go amqp.Listen()
	}

	// Run gRPC
	if GRPC_ENABLE == "true" {
		go grpc.Listen()
	}

	// Run Prometheus
	if PROMETHEUS_ENABLE == "true" {
		go metrics.Listen()
		go healthcheck.Listen()
	}

	// Wait forever
	select {}
}
