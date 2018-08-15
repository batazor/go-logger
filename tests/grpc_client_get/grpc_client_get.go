package main

import (
	"fmt"
	"github.com/batazor/go-logger/pb"
	"github.com/batazor/go-logger/utils"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	log       = logrus.New()
	GRPC_PORT = utils.Getenv("GRPC_PORT", "50051")
	client    telemetry.TelemetryClient
)

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)

	// Connect to InfluxDB
	port := fmt.Sprintf(":%s", GRPC_PORT)
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Open port: ", err)
	}

	client = telemetry.NewTelemetryClient(conn)
}

func main() {
	res, err := client.GetPacket(context.Background(), &telemetry.PacketRequest{
		Packet: "May",
	})
	if err != nil {
		log.Fatal("Error GetPacket: ", err)
	}

	log.Info("RESULT: ", res.Packet)
}
