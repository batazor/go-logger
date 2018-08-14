package main

import (
	"context"
	"github.com/batazor/go-logger/modules/amqp"
	"github.com/batazor/go-logger/modules/grpc"
	"github.com/batazor/go-logger/modules/influxdb"
	"github.com/batazor/go-logger/modules/jaeger"
	"github.com/batazor/go-logger/modules/metrics"
	"github.com/batazor/go-logger/utils"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

var (
	// Logging
	log = logrus.New()

	// Channel
	packetCh = make(chan []byte)

	// ENV
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
	// Run InfluxDB
	go func() {
		influxdb.Connect(packetCh)
	}()

	// Run AMQP
	if AMQP_ENABLE == "true" {
		go amqp.Listen(packetCh)
	}

	// Run AMQP
	if PROMETHEUS_ENABLE == "true" {
		go metrics.Listen()
	}

	// Run gRPC
	if GRPC_ENABLE == "true" {
		grpc.Listen()
	}
}
