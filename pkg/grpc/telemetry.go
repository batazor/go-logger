package grpc

import (
	"context"
	"github.com/batazor/go-logger/pb"
	"github.com/batazor/go-logger/pkg/redis"
)

func (s *server) GetPacket(ctx context.Context, in *telemetry.PacketRequest) (*telemetry.DataResponse, error) {
	//r := redis.Insert(in.Packet)
	//
	//return &telemetry.DataResponse{
	//	Packet: r,
	//}, nil
	return nil, nil
}

func (s *server) SendPacket(ctx context.Context, in *telemetry.PacketRequest) (*telemetry.PacketResponse, error) {
	r := redis.Insert([]byte(in.Packet))

	return &telemetry.PacketResponse{Success: r}, nil
}
