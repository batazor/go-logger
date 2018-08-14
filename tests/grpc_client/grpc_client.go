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
)

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)
}

func main() {
	port := fmt.Sprintf(":%s", GRPC_PORT)
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Open port: ", err)
	}
	defer conn.Close()

	client := telemetry.NewTelemetryClient(conn)
	res, err := client.GetPacket(context.Background(), &telemetry.PacketRequest{
		Packet: "July",
	})
	if err != nil {
		log.Fatal("Error GetPacket: ", err)
	}

	log.Info("RESULT: ", res.Packet)
}
