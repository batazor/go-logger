package grpc

import (
	"context"
	"github.com/batazor/go-logger/modules/influxdb"
	"github.com/batazor/go-logger/pb"
)

func (s *server) GetPacket(ctx context.Context, in *telemetry.PacketRequest) (*telemetry.DataResponse, error) {
	log.Info("in.Packet: ", in.Packet)

	r := influxdb.Query(in.Packet)

	return &telemetry.DataResponse{
		Packet: r,
	}, nil
}

func (s *server) SendPacket(ctx context.Context, in *telemetry.PacketRequest) (*telemetry.PacketResponse, error) {
	return &telemetry.PacketResponse{Success: true}, nil
}
