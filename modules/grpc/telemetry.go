package grpc

import (
	"context"
	"github.com/batazor/go-logger/modules/influxdb"
	"github.com/batazor/go-logger/pb"
)

func (s *server) GetPacket(ctx context.Context, in *telemetry.PacketRequest) (*telemetry.DataResponse, error) {
	log.Info("in.Packet: ", in.Packet)

	influxdb.Query(in.Packet)

	//log.Info("RESPONSE: ", r)

	return &telemetry.DataResponse{
		Packet: "",
	}, nil
}

func (s *server) SendPacket(ctx context.Context, in *telemetry.PacketRequest) (*telemetry.PacketResponse, error) {
	return &telemetry.PacketResponse{Success: true}, nil
}
