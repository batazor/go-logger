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
	log.Info("in.Packet: ", in.Packet)

	r := influxdb.InsertJSON(in.Packet)

	return &telemetry.PacketResponse{Success: r}, nil
}
