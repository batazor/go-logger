package grpc

import (
	"context"
	"github.com/batazor/go-logger/pb"
)

func (s *server) GetPacket(ctx context.Context, in *telemetry.PacketRequest) (*telemetry.DataResponse, error) {
	s.apiDBRequest <- Request{
		Data: in.Packet,
		Response: func(name string) {
			log.Info("RESPONSE: ", name)
		},
	}

	return &telemetry.DataResponse{
		Packet: "",
	}, nil
}

func (s *server) SendPacket(ctx context.Context, in *telemetry.PacketRequest) (*telemetry.PacketResponse, error) {
	return &telemetry.PacketResponse{Success: true}, nil
}
