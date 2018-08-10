package main

import (
	"github.com/batazor/go-logger/modules/amqp"
	"github.com/batazor/go-logger/modules/grpc"
	"github.com/batazor/go-logger/modules/influxdb"
	"github.com/batazor/go-logger/modules/metrics"
	"github.com/batazor/go-logger/utils"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()

	packetCh          = make(chan []byte)
	AMQP_ENABLE       = utils.Getenv("AMQP_ENABLE", "true")
	PROMETHEUS_ENABLE = utils.Getenv("PROMETHEUS_ENABLE", "true")
	GRPC_ENABLE       = utils.Getenv("GRPC_ENABLE", "true")
)

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)
}

func main() {
	// Run InfluxDB
	go influxdb.Connect(packetCh)

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
