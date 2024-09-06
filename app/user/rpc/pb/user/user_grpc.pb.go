// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.11
// source: user.proto

package user

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
	UserRegister_Register_FullMethodName = "/user.UserRegister/Register"
)

// UserRegisterClient is the client API for UserRegister service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserRegisterClient interface {
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRes, error)
}

type userRegisterClient struct {
	cc grpc.ClientConnInterface
}

func NewUserRegisterClient(cc grpc.ClientConnInterface) UserRegisterClient {
	return &userRegisterClient{cc}
}

func (c *userRegisterClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRes, error) {
	out := new(RegisterRes)
	err := c.cc.Invoke(ctx, UserRegister_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserRegisterServer is the server API for UserRegister service.
// All implementations must embed UnimplementedUserRegisterServer
// for forward compatibility
type UserRegisterServer interface {
	Register(context.Context, *RegisterReq) (*RegisterRes, error)
	mustEmbedUnimplementedUserRegisterServer()
}

// UnimplementedUserRegisterServer must be embedded to have forward compatible implementations.
type UnimplementedUserRegisterServer struct {
}

func (UnimplementedUserRegisterServer) Register(context.Context, *RegisterReq) (*RegisterRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserRegisterServer) mustEmbedUnimplementedUserRegisterServer() {}

// UnsafeUserRegisterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserRegisterServer will
// result in compilation errors.
type UnsafeUserRegisterServer interface {
	mustEmbedUnimplementedUserRegisterServer()
}

func RegisterUserRegisterServer(s grpc.ServiceRegistrar, srv UserRegisterServer) {
	s.RegisterService(&UserRegister_ServiceDesc, srv)
}

func _UserRegister_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegisterServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRegister_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegisterServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserRegister_ServiceDesc is the grpc.ServiceDesc for UserRegister service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserRegister_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserRegister",
	HandlerType: (*UserRegisterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _UserRegister_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
