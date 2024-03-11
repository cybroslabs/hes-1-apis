// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: driver.proto

package grpcapi

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

const (
	Driver_StartJob_FullMethodName  = "/grpcapi.Driver/StartJob"
	Driver_CancelJob_FullMethodName = "/grpcapi.Driver/CancelJob"
)

// DriverClient is the client API for Driver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DriverClient interface {
	StartJob(ctx context.Context, in *StartJobRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	CancelJob(ctx context.Context, in *CancelJobRequest, opts ...grpc.CallOption) (*CommonResponse, error)
}

type driverClient struct {
	cc grpc.ClientConnInterface
}

func NewDriverClient(cc grpc.ClientConnInterface) DriverClient {
	return &driverClient{cc}
}

func (c *driverClient) StartJob(ctx context.Context, in *StartJobRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, Driver_StartJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) CancelJob(ctx context.Context, in *CancelJobRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, Driver_CancelJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriverServer is the server API for Driver service.
// All implementations must embed UnimplementedDriverServer
// for forward compatibility
type DriverServer interface {
	StartJob(context.Context, *StartJobRequest) (*CommonResponse, error)
	CancelJob(context.Context, *CancelJobRequest) (*CommonResponse, error)
	mustEmbedUnimplementedDriverServer()
}

// UnimplementedDriverServer must be embedded to have forward compatible implementations.
type UnimplementedDriverServer struct {
}

func (UnimplementedDriverServer) StartJob(context.Context, *StartJobRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartJob not implemented")
}
func (UnimplementedDriverServer) CancelJob(context.Context, *CancelJobRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelJob not implemented")
}
func (UnimplementedDriverServer) mustEmbedUnimplementedDriverServer() {}

// UnsafeDriverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DriverServer will
// result in compilation errors.
type UnsafeDriverServer interface {
	mustEmbedUnimplementedDriverServer()
}

func RegisterDriverServer(s grpc.ServiceRegistrar, srv DriverServer) {
	s.RegisterService(&Driver_ServiceDesc, srv)
}

func _Driver_StartJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).StartJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Driver_StartJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).StartJob(ctx, req.(*StartJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_CancelJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).CancelJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Driver_CancelJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).CancelJob(ctx, req.(*CancelJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Driver_ServiceDesc is the grpc.ServiceDesc for Driver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Driver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcapi.Driver",
	HandlerType: (*DriverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartJob",
			Handler:    _Driver_StartJob_Handler,
		},
		{
			MethodName: "CancelJob",
			Handler:    _Driver_CancelJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "driver.proto",
}

const (
	DriverHost_NegotiateStart_FullMethodName = "/grpcapi.DriverHost/NegotiateStart"
	DriverHost_FinishActions_FullMethodName  = "/grpcapi.DriverHost/FinishActions"
	DriverHost_FinishJob_FullMethodName      = "/grpcapi.DriverHost/FinishJob"
	DriverHost_CacheSet_FullMethodName       = "/grpcapi.DriverHost/CacheSet"
	DriverHost_CacheGet_FullMethodName       = "/grpcapi.DriverHost/CacheGet"
)

// DriverHostClient is the client API for DriverHost service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DriverHostClient interface {
	NegotiateStart(ctx context.Context, in *NegotiateRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	FinishActions(ctx context.Context, in *FinishActionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	FinishJob(ctx context.Context, in *FinishJobRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CacheSet(ctx context.Context, in *CacheSetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CacheGet(ctx context.Context, in *CacheGetRequest, opts ...grpc.CallOption) (*CacheGetResponse, error)
}

type driverHostClient struct {
	cc grpc.ClientConnInterface
}

func NewDriverHostClient(cc grpc.ClientConnInterface) DriverHostClient {
	return &driverHostClient{cc}
}

func (c *driverHostClient) NegotiateStart(ctx context.Context, in *NegotiateRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, DriverHost_NegotiateStart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverHostClient) FinishActions(ctx context.Context, in *FinishActionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DriverHost_FinishActions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverHostClient) FinishJob(ctx context.Context, in *FinishJobRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DriverHost_FinishJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverHostClient) CacheSet(ctx context.Context, in *CacheSetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DriverHost_CacheSet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverHostClient) CacheGet(ctx context.Context, in *CacheGetRequest, opts ...grpc.CallOption) (*CacheGetResponse, error) {
	out := new(CacheGetResponse)
	err := c.cc.Invoke(ctx, DriverHost_CacheGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriverHostServer is the server API for DriverHost service.
// All implementations must embed UnimplementedDriverHostServer
// for forward compatibility
type DriverHostServer interface {
	NegotiateStart(context.Context, *NegotiateRequest) (*CommonResponse, error)
	FinishActions(context.Context, *FinishActionRequest) (*emptypb.Empty, error)
	FinishJob(context.Context, *FinishJobRequest) (*emptypb.Empty, error)
	CacheSet(context.Context, *CacheSetRequest) (*emptypb.Empty, error)
	CacheGet(context.Context, *CacheGetRequest) (*CacheGetResponse, error)
	mustEmbedUnimplementedDriverHostServer()
}

// UnimplementedDriverHostServer must be embedded to have forward compatible implementations.
type UnimplementedDriverHostServer struct {
}

func (UnimplementedDriverHostServer) NegotiateStart(context.Context, *NegotiateRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NegotiateStart not implemented")
}
func (UnimplementedDriverHostServer) FinishActions(context.Context, *FinishActionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishActions not implemented")
}
func (UnimplementedDriverHostServer) FinishJob(context.Context, *FinishJobRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishJob not implemented")
}
func (UnimplementedDriverHostServer) CacheSet(context.Context, *CacheSetRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CacheSet not implemented")
}
func (UnimplementedDriverHostServer) CacheGet(context.Context, *CacheGetRequest) (*CacheGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CacheGet not implemented")
}
func (UnimplementedDriverHostServer) mustEmbedUnimplementedDriverHostServer() {}

// UnsafeDriverHostServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DriverHostServer will
// result in compilation errors.
type UnsafeDriverHostServer interface {
	mustEmbedUnimplementedDriverHostServer()
}

func RegisterDriverHostServer(s grpc.ServiceRegistrar, srv DriverHostServer) {
	s.RegisterService(&DriverHost_ServiceDesc, srv)
}

func _DriverHost_NegotiateStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NegotiateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverHostServer).NegotiateStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverHost_NegotiateStart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverHostServer).NegotiateStart(ctx, req.(*NegotiateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverHost_FinishActions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverHostServer).FinishActions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverHost_FinishActions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverHostServer).FinishActions(ctx, req.(*FinishActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverHost_FinishJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverHostServer).FinishJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverHost_FinishJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverHostServer).FinishJob(ctx, req.(*FinishJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverHost_CacheSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CacheSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverHostServer).CacheSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverHost_CacheSet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverHostServer).CacheSet(ctx, req.(*CacheSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverHost_CacheGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CacheGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverHostServer).CacheGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverHost_CacheGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverHostServer).CacheGet(ctx, req.(*CacheGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DriverHost_ServiceDesc is the grpc.ServiceDesc for DriverHost service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DriverHost_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcapi.DriverHost",
	HandlerType: (*DriverHostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NegotiateStart",
			Handler:    _DriverHost_NegotiateStart_Handler,
		},
		{
			MethodName: "FinishActions",
			Handler:    _DriverHost_FinishActions_Handler,
		},
		{
			MethodName: "FinishJob",
			Handler:    _DriverHost_FinishJob_Handler,
		},
		{
			MethodName: "CacheSet",
			Handler:    _DriverHost_CacheSet_Handler,
		},
		{
			MethodName: "CacheGet",
			Handler:    _DriverHost_CacheGet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "driver.proto",
}
