package grpc

import (
	"github.com/batazor/go-logger/pb"
	"golang.org/x/net/context"
)

func (s *server) GetPacket(ctx context.Context, in *telemetry.PacketRequest) (*telemetry.DataResponse, error) {
	return &telemetry.DataResponse{
		Packet: "",
	}, nil
}

func (s *server) SendPacket(ctx context.Context, in *telemetry.PacketRequest) (*telemetry.PacketResponse, error) {
	return &telemetry.PacketResponse{Success: true}, nil
}
