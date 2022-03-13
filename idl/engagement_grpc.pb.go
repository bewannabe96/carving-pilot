// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: carving/engagement.proto

package carving

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

// EngagementServiceClient is the client API for EngagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EngagementServiceClient interface {
	LogScreenChange(ctx context.Context, in *LogScreenChangeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	LogEvent(ctx context.Context, in *LogEventRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type engagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEngagementServiceClient(cc grpc.ClientConnInterface) EngagementServiceClient {
	return &engagementServiceClient{cc}
}

func (c *engagementServiceClient) LogScreenChange(ctx context.Context, in *LogScreenChangeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/carving.EngagementService/LogScreenChange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engagementServiceClient) LogEvent(ctx context.Context, in *LogEventRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/carving.EngagementService/LogEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EngagementServiceServer is the server API for EngagementService service.
// All implementations must embed UnimplementedEngagementServiceServer
// for forward compatibility
type EngagementServiceServer interface {
	LogScreenChange(context.Context, *LogScreenChangeRequest) (*emptypb.Empty, error)
	LogEvent(context.Context, *LogEventRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedEngagementServiceServer()
}

// UnimplementedEngagementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEngagementServiceServer struct {
}

func (UnimplementedEngagementServiceServer) LogScreenChange(context.Context, *LogScreenChangeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogScreenChange not implemented")
}
func (UnimplementedEngagementServiceServer) LogEvent(context.Context, *LogEventRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogEvent not implemented")
}
func (UnimplementedEngagementServiceServer) mustEmbedUnimplementedEngagementServiceServer() {}

// UnsafeEngagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EngagementServiceServer will
// result in compilation errors.
type UnsafeEngagementServiceServer interface {
	mustEmbedUnimplementedEngagementServiceServer()
}

func RegisterEngagementServiceServer(s grpc.ServiceRegistrar, srv EngagementServiceServer) {
	s.RegisterService(&EngagementService_ServiceDesc, srv)
}

func _EngagementService_LogScreenChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogScreenChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngagementServiceServer).LogScreenChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/carving.EngagementService/LogScreenChange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngagementServiceServer).LogScreenChange(ctx, req.(*LogScreenChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EngagementService_LogEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngagementServiceServer).LogEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/carving.EngagementService/LogEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngagementServiceServer).LogEvent(ctx, req.(*LogEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EngagementService_ServiceDesc is the grpc.ServiceDesc for EngagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EngagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "carving.EngagementService",
	HandlerType: (*EngagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LogScreenChange",
			Handler:    _EngagementService_LogScreenChange_Handler,
		},
		{
			MethodName: "LogEvent",
			Handler:    _EngagementService_LogEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "carving/engagement.proto",
}