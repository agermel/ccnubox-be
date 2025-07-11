// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: static/v1/static.proto

package staticv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	StaticService_GetStaticByName_FullMethodName    = "/static.v1.StaticService/GetStaticByName"
	StaticService_SaveStatic_FullMethodName         = "/static.v1.StaticService/SaveStatic"
	StaticService_GetStaticsByLabels_FullMethodName = "/static.v1.StaticService/GetStaticsByLabels"
)

// StaticServiceClient is the client API for StaticService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StaticServiceClient interface {
	GetStaticByName(ctx context.Context, in *GetStaticByNameRequest, opts ...grpc.CallOption) (*GetStaticByNameResponse, error)
	SaveStatic(ctx context.Context, in *SaveStaticRequest, opts ...grpc.CallOption) (*SaveStaticResponse, error)
	GetStaticsByLabels(ctx context.Context, in *GetStaticsByLabelsRequest, opts ...grpc.CallOption) (*GetStaticsByLabelsResponse, error)
}

type staticServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStaticServiceClient(cc grpc.ClientConnInterface) StaticServiceClient {
	return &staticServiceClient{cc}
}

func (c *staticServiceClient) GetStaticByName(ctx context.Context, in *GetStaticByNameRequest, opts ...grpc.CallOption) (*GetStaticByNameResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStaticByNameResponse)
	err := c.cc.Invoke(ctx, StaticService_GetStaticByName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticServiceClient) SaveStatic(ctx context.Context, in *SaveStaticRequest, opts ...grpc.CallOption) (*SaveStaticResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SaveStaticResponse)
	err := c.cc.Invoke(ctx, StaticService_SaveStatic_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticServiceClient) GetStaticsByLabels(ctx context.Context, in *GetStaticsByLabelsRequest, opts ...grpc.CallOption) (*GetStaticsByLabelsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStaticsByLabelsResponse)
	err := c.cc.Invoke(ctx, StaticService_GetStaticsByLabels_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StaticServiceServer is the server API for StaticService service.
// All implementations must embed UnimplementedStaticServiceServer
// for forward compatibility.
type StaticServiceServer interface {
	GetStaticByName(context.Context, *GetStaticByNameRequest) (*GetStaticByNameResponse, error)
	SaveStatic(context.Context, *SaveStaticRequest) (*SaveStaticResponse, error)
	GetStaticsByLabels(context.Context, *GetStaticsByLabelsRequest) (*GetStaticsByLabelsResponse, error)
	mustEmbedUnimplementedStaticServiceServer()
}

// UnimplementedStaticServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStaticServiceServer struct{}

func (UnimplementedStaticServiceServer) GetStaticByName(context.Context, *GetStaticByNameRequest) (*GetStaticByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStaticByName not implemented")
}
func (UnimplementedStaticServiceServer) SaveStatic(context.Context, *SaveStaticRequest) (*SaveStaticResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveStatic not implemented")
}
func (UnimplementedStaticServiceServer) GetStaticsByLabels(context.Context, *GetStaticsByLabelsRequest) (*GetStaticsByLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStaticsByLabels not implemented")
}
func (UnimplementedStaticServiceServer) mustEmbedUnimplementedStaticServiceServer() {}
func (UnimplementedStaticServiceServer) testEmbeddedByValue()                       {}

// UnsafeStaticServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StaticServiceServer will
// result in compilation errors.
type UnsafeStaticServiceServer interface {
	mustEmbedUnimplementedStaticServiceServer()
}

func RegisterStaticServiceServer(s grpc.ServiceRegistrar, srv StaticServiceServer) {
	// If the following call pancis, it indicates UnimplementedStaticServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StaticService_ServiceDesc, srv)
}

func _StaticService_GetStaticByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStaticByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticServiceServer).GetStaticByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaticService_GetStaticByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticServiceServer).GetStaticByName(ctx, req.(*GetStaticByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaticService_SaveStatic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveStaticRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticServiceServer).SaveStatic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaticService_SaveStatic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticServiceServer).SaveStatic(ctx, req.(*SaveStaticRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaticService_GetStaticsByLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStaticsByLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticServiceServer).GetStaticsByLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaticService_GetStaticsByLabels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticServiceServer).GetStaticsByLabels(ctx, req.(*GetStaticsByLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StaticService_ServiceDesc is the grpc.ServiceDesc for StaticService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StaticService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "static.v1.StaticService",
	HandlerType: (*StaticServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStaticByName",
			Handler:    _StaticService_GetStaticByName_Handler,
		},
		{
			MethodName: "SaveStatic",
			Handler:    _StaticService_SaveStatic_Handler,
		},
		{
			MethodName: "GetStaticsByLabels",
			Handler:    _StaticService_GetStaticsByLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "static/v1/static.proto",
}
