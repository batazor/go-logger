package grpc

import (
	"fmt"
	"github.com/batazor/go-logger/pb"
	"github.com/batazor/go-logger/utils"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

var (
	log       = logrus.New()
	GRPC_PORT = utils.Getenv("GRPC_PORT", "50051")
)

type server struct{}

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)
}

// Connect to API by gRPC
func Listen() {
	port := fmt.Sprintf(":%s", GRPC_PORT)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Open port: ", err)
	}

	s := grpc.NewServer()
	pb.RegisterTelemetryServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	} else {
		log.Info("Run gRPC on port " + port)
	}

	log.Info("Success connect to Mail API by gRPC")
}
