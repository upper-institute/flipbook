// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/v1/eventhandler.proto

package flipbookv1

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

// EventHandlerClient is the client API for EventHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventHandlerClient interface {
	Handle(ctx context.Context, opts ...grpc.CallOption) (EventHandler_HandleClient, error)
}

type eventHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewEventHandlerClient(cc grpc.ClientConnInterface) EventHandlerClient {
	return &eventHandlerClient{cc}
}

func (c *eventHandlerClient) Handle(ctx context.Context, opts ...grpc.CallOption) (EventHandler_HandleClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventHandler_ServiceDesc.Streams[0], "/flipbook.v1.EventHandler/Handle", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventHandlerHandleClient{stream}
	return x, nil
}

type EventHandler_HandleClient interface {
	Send(*Event) error
	Recv() (*Commit, error)
	grpc.ClientStream
}

type eventHandlerHandleClient struct {
	grpc.ClientStream
}

func (x *eventHandlerHandleClient) Send(m *Event) error {
	return x.ClientStream.SendMsg(m)
}

func (x *eventHandlerHandleClient) Recv() (*Commit, error) {
	m := new(Commit)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventHandlerServer is the server API for EventHandler service.
// All implementations must embed UnimplementedEventHandlerServer
// for forward compatibility
type EventHandlerServer interface {
	Handle(EventHandler_HandleServer) error
	mustEmbedUnimplementedEventHandlerServer()
}

// UnimplementedEventHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedEventHandlerServer struct {
}

func (UnimplementedEventHandlerServer) Handle(EventHandler_HandleServer) error {
	return status.Errorf(codes.Unimplemented, "method Handle not implemented")
}
func (UnimplementedEventHandlerServer) mustEmbedUnimplementedEventHandlerServer() {}

// UnsafeEventHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventHandlerServer will
// result in compilation errors.
type UnsafeEventHandlerServer interface {
	mustEmbedUnimplementedEventHandlerServer()
}

func RegisterEventHandlerServer(s grpc.ServiceRegistrar, srv EventHandlerServer) {
	s.RegisterService(&EventHandler_ServiceDesc, srv)
}

func _EventHandler_Handle_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EventHandlerServer).Handle(&eventHandlerHandleServer{stream})
}

type EventHandler_HandleServer interface {
	Send(*Commit) error
	Recv() (*Event, error)
	grpc.ServerStream
}

type eventHandlerHandleServer struct {
	grpc.ServerStream
}

func (x *eventHandlerHandleServer) Send(m *Commit) error {
	return x.ServerStream.SendMsg(m)
}

func (x *eventHandlerHandleServer) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventHandler_ServiceDesc is the grpc.ServiceDesc for EventHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flipbook.v1.EventHandler",
	HandlerType: (*EventHandlerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Handle",
			Handler:       _EventHandler_Handle_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/v1/eventhandler.proto",
}
