package grpc

import (
	"fmt"
	probe "github.com/batazor/go-logger/modules/healthcheck"
	"github.com/batazor/go-logger/pb"
	"github.com/batazor/go-logger/utils"
	"github.com/heptiolabs/healthcheck"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"time"
)

var (
	log       = logrus.New()
	GRPC_PORT = utils.Getenv("GRPC_PORT", "50051")
)

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)
}

type server struct{}

func Listen() {
	port := fmt.Sprintf(":%s", GRPC_PORT)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Open port: ", err)
	}

	// Health check
	probe.Health.AddReadinessCheck(
		"grpc",
		healthcheck.Timeout(func() error { return err }, time.Second*10))

	log.Info("Run gRPC on port " + port)

	grpcServer := grpc.NewServer()
	telemetry.RegisterTelemetryServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
