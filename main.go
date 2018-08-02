package main

import (
	"fmt"
	"github.com/batazor/go-logger/modules/amqp"
	"github.com/batazor/go-logger/modules/influxdb"
	"github.com/batazor/go-logger/modules/telemetry"
	"github.com/batazor/go-logger/utils"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
)

var (
	log = logrus.New()

	packetCh    = make(chan []byte)
	AMQP_ENABLE = utils.Getenv("AMQP_ENABLE", "false")
	GRPC_ENABLE = utils.Getenv("GRPC_ENABLE", "true")
	GRPC_PORT   = utils.Getenv("GRPC_PORT", "50051")
)

type server struct{}

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)
}

func (s *server) SendPacket(ctx context.Context, in *telemetry.PacketRequest) (*telemetry.PacketResponse, error) {
	return &telemetry.PacketResponse{Success: true}, nil
}

func main() {
	// Run InfluxDB
	go influxdb.Connect(packetCh)

	// Run AMQP
	if AMQP_ENABLE == "true" {
		go amqp.Listen(packetCh)
	}

	// Run gRPC
	if GRPC_ENABLE == "true" {
		port := fmt.Sprintf(":%s", GRPC_PORT)
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatal("Open port: ", err)
		}
		s := grpc.NewServer()
		telemetry.RegisterTelemetryServer(s, &server{})
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to server: %v", err)
		} else {
			log.Info("Run gRPC on port " + port)
		}
	}
}
