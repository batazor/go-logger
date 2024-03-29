// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/telemetry.proto

package telemetry

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PacketRequest struct {
	Packet               string   `protobuf:"bytes,1,opt,name=Packet,proto3" json:"Packet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PacketRequest) Reset()         { *m = PacketRequest{} }
func (m *PacketRequest) String() string { return proto.CompactTextString(m) }
func (*PacketRequest) ProtoMessage()    {}
func (*PacketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_telemetry_016e21ca5566e7cb, []int{0}
}
func (m *PacketRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PacketRequest.Unmarshal(m, b)
}
func (m *PacketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PacketRequest.Marshal(b, m, deterministic)
}
func (dst *PacketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketRequest.Merge(dst, src)
}
func (m *PacketRequest) XXX_Size() int {
	return xxx_messageInfo_PacketRequest.Size(m)
}
func (m *PacketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PacketRequest proto.InternalMessageInfo

func (m *PacketRequest) GetPacket() string {
	if m != nil {
		return m.Packet
	}
	return ""
}

type PacketResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PacketResponse) Reset()         { *m = PacketResponse{} }
func (m *PacketResponse) String() string { return proto.CompactTextString(m) }
func (*PacketResponse) ProtoMessage()    {}
func (*PacketResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_telemetry_016e21ca5566e7cb, []int{1}
}
func (m *PacketResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PacketResponse.Unmarshal(m, b)
}
func (m *PacketResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PacketResponse.Marshal(b, m, deterministic)
}
func (dst *PacketResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketResponse.Merge(dst, src)
}
func (m *PacketResponse) XXX_Size() int {
	return xxx_messageInfo_PacketResponse.Size(m)
}
func (m *PacketResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PacketResponse proto.InternalMessageInfo

func (m *PacketResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type DataResponse struct {
	Packet               []byte   `protobuf:"bytes,1,opt,name=Packet,proto3" json:"Packet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataResponse) Reset()         { *m = DataResponse{} }
func (m *DataResponse) String() string { return proto.CompactTextString(m) }
func (*DataResponse) ProtoMessage()    {}
func (*DataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_telemetry_016e21ca5566e7cb, []int{2}
}
func (m *DataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataResponse.Unmarshal(m, b)
}
func (m *DataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataResponse.Marshal(b, m, deterministic)
}
func (dst *DataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataResponse.Merge(dst, src)
}
func (m *DataResponse) XXX_Size() int {
	return xxx_messageInfo_DataResponse.Size(m)
}
func (m *DataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DataResponse proto.InternalMessageInfo

func (m *DataResponse) GetPacket() []byte {
	if m != nil {
		return m.Packet
	}
	return nil
}

func init() {
	proto.RegisterType((*PacketRequest)(nil), "telemetry.PacketRequest")
	proto.RegisterType((*PacketResponse)(nil), "telemetry.PacketResponse")
	proto.RegisterType((*DataResponse)(nil), "telemetry.DataResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TelemetryClient is the client API for Telemetry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TelemetryClient interface {
	SendPacket(ctx context.Context, in *PacketRequest, opts ...grpc.CallOption) (*PacketResponse, error)
	GetPacket(ctx context.Context, in *PacketRequest, opts ...grpc.CallOption) (*DataResponse, error)
}

type telemetryClient struct {
	cc *grpc.ClientConn
}

func NewTelemetryClient(cc *grpc.ClientConn) TelemetryClient {
	return &telemetryClient{cc}
}

func (c *telemetryClient) SendPacket(ctx context.Context, in *PacketRequest, opts ...grpc.CallOption) (*PacketResponse, error) {
	out := new(PacketResponse)
	err := c.cc.Invoke(ctx, "/telemetry.Telemetry/SendPacket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryClient) GetPacket(ctx context.Context, in *PacketRequest, opts ...grpc.CallOption) (*DataResponse, error) {
	out := new(DataResponse)
	err := c.cc.Invoke(ctx, "/telemetry.Telemetry/GetPacket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TelemetryServer is the server API for Telemetry service.
type TelemetryServer interface {
	SendPacket(context.Context, *PacketRequest) (*PacketResponse, error)
	GetPacket(context.Context, *PacketRequest) (*DataResponse, error)
}

func RegisterTelemetryServer(s *grpc.Server, srv TelemetryServer) {
	s.RegisterService(&_Telemetry_serviceDesc, srv)
}

func _Telemetry_SendPacket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PacketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServer).SendPacket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telemetry.Telemetry/SendPacket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServer).SendPacket(ctx, req.(*PacketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Telemetry_GetPacket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PacketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServer).GetPacket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telemetry.Telemetry/GetPacket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServer).GetPacket(ctx, req.(*PacketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Telemetry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "telemetry.Telemetry",
	HandlerType: (*TelemetryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendPacket",
			Handler:    _Telemetry_SendPacket_Handler,
		},
		{
			MethodName: "GetPacket",
			Handler:    _Telemetry_GetPacket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/telemetry.proto",
}

func init() { proto.RegisterFile("pb/telemetry.proto", fileDescriptor_telemetry_016e21ca5566e7cb) }

var fileDescriptor_telemetry_016e21ca5566e7cb = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x48, 0xd2, 0x2f,
	0x49, 0xcd, 0x49, 0xcd, 0x4d, 0x2d, 0x29, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x84, 0x0b, 0x28, 0xa9, 0x73, 0xf1, 0x06, 0x24, 0x26, 0x67, 0xa7, 0x96, 0x04, 0xa5, 0x16, 0x96,
	0xa6, 0x16, 0x97, 0x08, 0x89, 0x71, 0xb1, 0x41, 0x04, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0xa0, 0x3c, 0x25, 0x2d, 0x2e, 0x3e, 0x98, 0xc2, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x09,
	0x2e, 0xf6, 0xe2, 0xd2, 0xe4, 0xe4, 0xd4, 0xe2, 0x62, 0xb0, 0x52, 0x8e, 0x20, 0x18, 0x57, 0x49,
	0x8d, 0x8b, 0xc7, 0x25, 0xb1, 0x24, 0x11, 0xae, 0x12, 0xd5, 0x4c, 0x1e, 0x98, 0x99, 0x46, 0x93,
	0x18, 0xb9, 0x38, 0x43, 0x60, 0x4e, 0x11, 0x72, 0xe6, 0xe2, 0x0a, 0x4e, 0xcd, 0x4b, 0x81, 0xc8,
	0x09, 0x49, 0xe8, 0x21, 0x5c, 0x8d, 0xe2, 0x42, 0x29, 0x49, 0x2c, 0x32, 0x10, 0x8b, 0x94, 0x18,
	0x84, 0x1c, 0xb8, 0x38, 0xdd, 0x53, 0x4b, 0x08, 0x9a, 0x21, 0x8e, 0x24, 0x83, 0xec, 0x54, 0x25,
	0x86, 0x24, 0x36, 0x70, 0x18, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x62, 0x1a, 0xff, 0x75,
	0x39, 0x01, 0x00, 0x00,
}
