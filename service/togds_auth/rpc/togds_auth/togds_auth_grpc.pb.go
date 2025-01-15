// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: togds_auth.proto

package togds_auth

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
	TogdsAuthRpc_CheckToken_FullMethodName = "/togds_auth.togdsAuthRpc/CheckToken"
	TogdsAuthRpc_Refresh_FullMethodName    = "/togds_auth.togdsAuthRpc/Refresh"
)

// TogdsAuthRpcClient is the client API for TogdsAuthRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TogdsAuthRpcClient interface {
	CheckToken(ctx context.Context, in *CheckTokenReq, opts ...grpc.CallOption) (*CheckTokenResp, error)
	Refresh(ctx context.Context, in *RefreshReq, opts ...grpc.CallOption) (*RefreshResp, error)
}

type togdsAuthRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewTogdsAuthRpcClient(cc grpc.ClientConnInterface) TogdsAuthRpcClient {
	return &togdsAuthRpcClient{cc}
}

func (c *togdsAuthRpcClient) CheckToken(ctx context.Context, in *CheckTokenReq, opts ...grpc.CallOption) (*CheckTokenResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckTokenResp)
	err := c.cc.Invoke(ctx, TogdsAuthRpc_CheckToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *togdsAuthRpcClient) Refresh(ctx context.Context, in *RefreshReq, opts ...grpc.CallOption) (*RefreshResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RefreshResp)
	err := c.cc.Invoke(ctx, TogdsAuthRpc_Refresh_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TogdsAuthRpcServer is the server API for TogdsAuthRpc service.
// All implementations must embed UnimplementedTogdsAuthRpcServer
// for forward compatibility.
type TogdsAuthRpcServer interface {
	CheckToken(context.Context, *CheckTokenReq) (*CheckTokenResp, error)
	Refresh(context.Context, *RefreshReq) (*RefreshResp, error)
	mustEmbedUnimplementedTogdsAuthRpcServer()
}

// UnimplementedTogdsAuthRpcServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTogdsAuthRpcServer struct{}

func (UnimplementedTogdsAuthRpcServer) CheckToken(context.Context, *CheckTokenReq) (*CheckTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckToken not implemented")
}
func (UnimplementedTogdsAuthRpcServer) Refresh(context.Context, *RefreshReq) (*RefreshResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedTogdsAuthRpcServer) mustEmbedUnimplementedTogdsAuthRpcServer() {}
func (UnimplementedTogdsAuthRpcServer) testEmbeddedByValue()                      {}

// UnsafeTogdsAuthRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TogdsAuthRpcServer will
// result in compilation errors.
type UnsafeTogdsAuthRpcServer interface {
	mustEmbedUnimplementedTogdsAuthRpcServer()
}

func RegisterTogdsAuthRpcServer(s grpc.ServiceRegistrar, srv TogdsAuthRpcServer) {
	// If the following call pancis, it indicates UnimplementedTogdsAuthRpcServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TogdsAuthRpc_ServiceDesc, srv)
}

func _TogdsAuthRpc_CheckToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TogdsAuthRpcServer).CheckToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TogdsAuthRpc_CheckToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TogdsAuthRpcServer).CheckToken(ctx, req.(*CheckTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TogdsAuthRpc_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TogdsAuthRpcServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TogdsAuthRpc_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TogdsAuthRpcServer).Refresh(ctx, req.(*RefreshReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TogdsAuthRpc_ServiceDesc is the grpc.ServiceDesc for TogdsAuthRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TogdsAuthRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "togds_auth.togdsAuthRpc",
	HandlerType: (*TogdsAuthRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckToken",
			Handler:    _TogdsAuthRpc_CheckToken_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _TogdsAuthRpc_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "togds_auth.proto",
}
