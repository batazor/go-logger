package grpc

import (
	"context"
	"github.com/batazor/go-logger/pb"
	"testing"
)

func TestGetPacket(t *testing.T) {
	s := server{}

	// set up test cases
	tests := []Request{
		{
			Data: "december",
			Response: func(name string) {
				log.Info("RESPONSE: ", name)
			},
		},
		{
			Data: "april",
			Response: func(name string) {
				log.Info("RESPONSE: ", name)
			},
		},
	}

	for _, tt := range tests {
		req := &telemetry.PacketRequest{Packet: tt.Data}
		resp, err := s.GetPacket(context.Background(), req)
		log.Info("CHECK", resp.Packet)
		if err != nil {
			t.Errorf("PacketRequest(%v) got unexpected error", err)
		}
		if resp.Packet != "" {
			t.Errorf("PacketRequest(%v)=%v, wanted %v", tt.Data, resp.Packet, true)
		}
	}
}
