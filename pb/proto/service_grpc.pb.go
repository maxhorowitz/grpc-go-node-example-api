// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: proto/service.proto

package proto

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Messenger_Send_FullMethodName = "/messenger.Messenger/Send"
)

// MessengerClient is the client API for Messenger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessengerClient interface {
	Send(ctx context.Context, opts ...grpc.CallOption) (Messenger_SendClient, error)
}

type messengerClient struct {
	cc grpc.ClientConnInterface
}

func NewMessengerClient(cc grpc.ClientConnInterface) MessengerClient {
	return &messengerClient{cc}
}

func (c *messengerClient) Send(ctx context.Context, opts ...grpc.CallOption) (Messenger_SendClient, error) {
	stream, err := c.cc.NewStream(ctx, &Messenger_ServiceDesc.Streams[0], Messenger_Send_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &messengerSendClient{stream}
	return x, nil
}

type Messenger_SendClient interface {
	Send(*Req) error
	Recv() (*Res, error)
	grpc.ClientStream
}

type messengerSendClient struct {
	grpc.ClientStream
}

func (x *messengerSendClient) Send(m *Req) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messengerSendClient) Recv() (*Res, error) {
	m := new(Res)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessengerServer is the server API for Messenger service.
// All implementations must embed UnimplementedMessengerServer
// for forward compatibility
type MessengerServer interface {
	Send(Messenger_SendServer) error
	mustEmbedUnimplementedMessengerServer()
}

// UnimplementedMessengerServer must be embedded to have forward compatible implementations.
type UnimplementedMessengerServer struct {
}

func (UnimplementedMessengerServer) Send(Messenger_SendServer) error {
	return status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedMessengerServer) mustEmbedUnimplementedMessengerServer() {}

// UnsafeMessengerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessengerServer will
// result in compilation errors.
type UnsafeMessengerServer interface {
	mustEmbedUnimplementedMessengerServer()
}

func RegisterMessengerServer(s grpc.ServiceRegistrar, srv MessengerServer) {
	s.RegisterService(&Messenger_ServiceDesc, srv)
}

func _Messenger_Send_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessengerServer).Send(&messengerSendServer{stream})
}

type Messenger_SendServer interface {
	Send(*Res) error
	Recv() (*Req, error)
	grpc.ServerStream
}

type messengerSendServer struct {
	grpc.ServerStream
}

func (x *messengerSendServer) Send(m *Res) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messengerSendServer) Recv() (*Req, error) {
	m := new(Req)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Messenger_ServiceDesc is the grpc.ServiceDesc for Messenger service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Messenger_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "messenger.Messenger",
	HandlerType: (*MessengerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Send",
			Handler:       _Messenger_Send_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/service.proto",
}
