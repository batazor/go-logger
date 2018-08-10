package grpc

import (
	"context"
	"github.com/batazor/go-logger/pb"
)

func (s *server) GetPacket(ctx context.Context, in *pb.PacketRequest) (*pb.DataResponse, error) {
	return &pb.DataResponse{
		Packet: "",
	}, nil
}

func (s *server) SendPacket(ctx context.Context, in *pb.PacketRequest) (*pb.PacketResponse, error) {
	return &pb.PacketResponse{Success: true}, nil
}
