// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: togdsuserrpc.proto

package togdsuserrpc

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
	AuthRpcre_Login_FullMethodName = "/togdsuserrpc.authRpcre/Login"
)

// AuthRpcreClient is the client API for AuthRpcre service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthRpcreClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
}

type authRpcreClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthRpcreClient(cc grpc.ClientConnInterface) AuthRpcreClient {
	return &authRpcreClient{cc}
}

func (c *authRpcreClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, AuthRpcre_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthRpcreServer is the server API for AuthRpcre service.
// All implementations must embed UnimplementedAuthRpcreServer
// for forward compatibility.
type AuthRpcreServer interface {
	Login(context.Context, *LoginReq) (*LoginResp, error)
	mustEmbedUnimplementedAuthRpcreServer()
}

// UnimplementedAuthRpcreServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthRpcreServer struct{}

func (UnimplementedAuthRpcreServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthRpcreServer) mustEmbedUnimplementedAuthRpcreServer() {}
func (UnimplementedAuthRpcreServer) testEmbeddedByValue()                   {}

// UnsafeAuthRpcreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthRpcreServer will
// result in compilation errors.
type UnsafeAuthRpcreServer interface {
	mustEmbedUnimplementedAuthRpcreServer()
}

func RegisterAuthRpcreServer(s grpc.ServiceRegistrar, srv AuthRpcreServer) {
	// If the following call pancis, it indicates UnimplementedAuthRpcreServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuthRpcre_ServiceDesc, srv)
}

func _AuthRpcre_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthRpcreServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthRpcre_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthRpcreServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthRpcre_ServiceDesc is the grpc.ServiceDesc for AuthRpcre service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthRpcre_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "togdsuserrpc.authRpcre",
	HandlerType: (*AuthRpcreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthRpcre_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "togdsuserrpc.proto",
}
