// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/v1/eventbus.proto

package flipbookv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EventBusClient is the client API for EventBus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventBusClient interface {
	Open(ctx context.Context, in *Channel_OpenRequest, opts ...grpc.CallOption) (*Channel, error)
	Close(ctx context.Context, in *Channel_CloseRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Watch(ctx context.Context, in *Channel_WatchRequest, opts ...grpc.CallOption) (EventBus_WatchClient, error)
}

type eventBusClient struct {
	cc grpc.ClientConnInterface
}

func NewEventBusClient(cc grpc.ClientConnInterface) EventBusClient {
	return &eventBusClient{cc}
}

func (c *eventBusClient) Open(ctx context.Context, in *Channel_OpenRequest, opts ...grpc.CallOption) (*Channel, error) {
	out := new(Channel)
	err := c.cc.Invoke(ctx, "/flipbook.v1.EventBus/Open", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventBusClient) Close(ctx context.Context, in *Channel_CloseRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/flipbook.v1.EventBus/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventBusClient) Watch(ctx context.Context, in *Channel_WatchRequest, opts ...grpc.CallOption) (EventBus_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventBus_ServiceDesc.Streams[0], "/flipbook.v1.EventBus/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventBusWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventBus_WatchClient interface {
	Recv() (*Channel, error)
	grpc.ClientStream
}

type eventBusWatchClient struct {
	grpc.ClientStream
}

func (x *eventBusWatchClient) Recv() (*Channel, error) {
	m := new(Channel)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventBusServer is the server API for EventBus service.
// All implementations must embed UnimplementedEventBusServer
// for forward compatibility
type EventBusServer interface {
	Open(context.Context, *Channel_OpenRequest) (*Channel, error)
	Close(context.Context, *Channel_CloseRequest) (*emptypb.Empty, error)
	Watch(*Channel_WatchRequest, EventBus_WatchServer) error
	mustEmbedUnimplementedEventBusServer()
}

// UnimplementedEventBusServer must be embedded to have forward compatible implementations.
type UnimplementedEventBusServer struct {
}

func (UnimplementedEventBusServer) Open(context.Context, *Channel_OpenRequest) (*Channel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Open not implemented")
}
func (UnimplementedEventBusServer) Close(context.Context, *Channel_CloseRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}
func (UnimplementedEventBusServer) Watch(*Channel_WatchRequest, EventBus_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (UnimplementedEventBusServer) mustEmbedUnimplementedEventBusServer() {}

// UnsafeEventBusServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventBusServer will
// result in compilation errors.
type UnsafeEventBusServer interface {
	mustEmbedUnimplementedEventBusServer()
}

func RegisterEventBusServer(s grpc.ServiceRegistrar, srv EventBusServer) {
	s.RegisterService(&EventBus_ServiceDesc, srv)
}

func _EventBus_Open_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Channel_OpenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventBusServer).Open(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flipbook.v1.EventBus/Open",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventBusServer).Open(ctx, req.(*Channel_OpenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventBus_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Channel_CloseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventBusServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flipbook.v1.EventBus/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventBusServer).Close(ctx, req.(*Channel_CloseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventBus_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Channel_WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventBusServer).Watch(m, &eventBusWatchServer{stream})
}

type EventBus_WatchServer interface {
	Send(*Channel) error
	grpc.ServerStream
}

type eventBusWatchServer struct {
	grpc.ServerStream
}

func (x *eventBusWatchServer) Send(m *Channel) error {
	return x.ServerStream.SendMsg(m)
}

// EventBus_ServiceDesc is the grpc.ServiceDesc for EventBus service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventBus_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flipbook.v1.EventBus",
	HandlerType: (*EventBusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Open",
			Handler:    _EventBus_Open_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _EventBus_Close_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _EventBus_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/v1/eventbus.proto",
}
