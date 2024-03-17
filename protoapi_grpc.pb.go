// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: protoapi.proto

package protoapi

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
	TimeChanger_GetParams_FullMethodName = "/TimeChanger/GetParams"
)

// TimeChangerClient is the client API for TimeChanger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimeChangerClient interface {
	GetParams(ctx context.Context, in *Params, opts ...grpc.CallOption) (*ChangedParams, error)
}

type timeChangerClient struct {
	cc grpc.ClientConnInterface
}

func NewTimeChangerClient(cc grpc.ClientConnInterface) TimeChangerClient {
	return &timeChangerClient{cc}
}

func (c *timeChangerClient) GetParams(ctx context.Context, in *Params, opts ...grpc.CallOption) (*ChangedParams, error) {
	out := new(ChangedParams)
	err := c.cc.Invoke(ctx, TimeChanger_GetParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimeChangerServer is the server API for TimeChanger service.
// All implementations must embed UnimplementedTimeChangerServer
// for forward compatibility
type TimeChangerServer interface {
	GetParams(context.Context, *Params) (*ChangedParams, error)
	mustEmbedUnimplementedTimeChangerServer()
}

// UnimplementedTimeChangerServer must be embedded to have forward compatible implementations.
type UnimplementedTimeChangerServer struct {
}

func (UnimplementedTimeChangerServer) GetParams(context.Context, *Params) (*ChangedParams, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParams not implemented")
}
func (UnimplementedTimeChangerServer) mustEmbedUnimplementedTimeChangerServer() {}

// UnsafeTimeChangerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimeChangerServer will
// result in compilation errors.
type UnsafeTimeChangerServer interface {
	mustEmbedUnimplementedTimeChangerServer()
}

func RegisterTimeChangerServer(s grpc.ServiceRegistrar, srv TimeChangerServer) {
	s.RegisterService(&TimeChanger_ServiceDesc, srv)
}

func _TimeChanger_GetParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Params)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeChangerServer).GetParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeChanger_GetParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeChangerServer).GetParams(ctx, req.(*Params))
	}
	return interceptor(ctx, in, info, handler)
}

// TimeChanger_ServiceDesc is the grpc.ServiceDesc for TimeChanger service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TimeChanger_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TimeChanger",
	HandlerType: (*TimeChangerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetParams",
			Handler:    _TimeChanger_GetParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protoapi.proto",
}
