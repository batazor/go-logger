package main

import (
	"encoding/json"
	"fmt"
	"github.com/batazor/go-logger/pb"
	"github.com/batazor/go-logger/utils"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
)

var (
	log       = logrus.New()
	GRPC_PORT = utils.Getenv("GRPC_PORT", "50051")
	client    telemetry.TelemetryClient
	INDEX     = 1
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
	packetCh := make(chan interface{}, 1)
	var task = func() {
		time.Sleep(time.Millisecond * 100)
		packet, _ := utils.GetRandomPacket()
		packetCh <- packet
	}
	go task()

	for {
		select {
		case res := <-packetCh:
			json, _ := json.Marshal(res)

			for i := 0; i < 100; i++ {
				go func() {
					res, err := client.SendPacket(context.Background(), &telemetry.PacketRequest{
						Packet: string(json),
					})

					if err != nil {
						log.Fatal("Error SendPacket: ", err)
					}

					INDEX += 1

					log.Info("RESULT ( ", INDEX, " ): ", res.Success)
				}()
			}

			go task()
		}
	}
}
