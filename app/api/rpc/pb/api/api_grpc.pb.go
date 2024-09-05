// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.11
// source: api.proto

package api

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
	UserService_Login_FullMethodName = "/api.UserService/Login"
	UserService_Info_FullMethodName  = "/api.UserService/Info"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	Login(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginRes, error)
	Info(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserLoginRes, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Login(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginRes, error) {
	out := new(UserLoginRes)
	err := c.cc.Invoke(ctx, UserService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Info(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserLoginRes, error) {
	out := new(UserLoginRes)
	err := c.cc.Invoke(ctx, UserService_Info_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	Login(context.Context, *UserLoginReq) (*UserLoginRes, error)
	Info(context.Context, *UserInfoReq) (*UserLoginRes, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Login(context.Context, *UserLoginReq) (*UserLoginRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServiceServer) Info(context.Context, *UserInfoReq) (*UserLoginRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*UserLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Info_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Info(ctx, req.(*UserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _UserService_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

const (
	StoreService_List_FullMethodName     = "/api.StoreService/List"
	StoreService_Info_FullMethodName     = "/api.StoreService/Info"
	StoreService_UserList_FullMethodName = "/api.StoreService/UserList"
)

// StoreServiceClient is the client API for StoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoreServiceClient interface {
	List(ctx context.Context, in *StoreListReq, opts ...grpc.CallOption) (*StoreListRes, error)
	Info(ctx context.Context, in *StoreInfoReq, opts ...grpc.CallOption) (*StoreInfoRes, error)
	UserList(ctx context.Context, in *StoreUsersReq, opts ...grpc.CallOption) (*StoreUsersRes, error)
}

type storeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreServiceClient(cc grpc.ClientConnInterface) StoreServiceClient {
	return &storeServiceClient{cc}
}

func (c *storeServiceClient) List(ctx context.Context, in *StoreListReq, opts ...grpc.CallOption) (*StoreListRes, error) {
	out := new(StoreListRes)
	err := c.cc.Invoke(ctx, StoreService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) Info(ctx context.Context, in *StoreInfoReq, opts ...grpc.CallOption) (*StoreInfoRes, error) {
	out := new(StoreInfoRes)
	err := c.cc.Invoke(ctx, StoreService_Info_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) UserList(ctx context.Context, in *StoreUsersReq, opts ...grpc.CallOption) (*StoreUsersRes, error) {
	out := new(StoreUsersRes)
	err := c.cc.Invoke(ctx, StoreService_UserList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServiceServer is the server API for StoreService service.
// All implementations must embed UnimplementedStoreServiceServer
// for forward compatibility
type StoreServiceServer interface {
	List(context.Context, *StoreListReq) (*StoreListRes, error)
	Info(context.Context, *StoreInfoReq) (*StoreInfoRes, error)
	UserList(context.Context, *StoreUsersReq) (*StoreUsersRes, error)
	mustEmbedUnimplementedStoreServiceServer()
}

// UnimplementedStoreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStoreServiceServer struct {
}

func (UnimplementedStoreServiceServer) List(context.Context, *StoreListReq) (*StoreListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedStoreServiceServer) Info(context.Context, *StoreInfoReq) (*StoreInfoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedStoreServiceServer) UserList(context.Context, *StoreUsersReq) (*StoreUsersRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserList not implemented")
}
func (UnimplementedStoreServiceServer) mustEmbedUnimplementedStoreServiceServer() {}

// UnsafeStoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoreServiceServer will
// result in compilation errors.
type UnsafeStoreServiceServer interface {
	mustEmbedUnimplementedStoreServiceServer()
}

func RegisterStoreServiceServer(s grpc.ServiceRegistrar, srv StoreServiceServer) {
	s.RegisterService(&StoreService_ServiceDesc, srv)
}

func _StoreService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).List(ctx, req.(*StoreListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreService_Info_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).Info(ctx, req.(*StoreInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_UserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreUsersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).UserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreService_UserList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).UserList(ctx, req.(*StoreUsersReq))
	}
	return interceptor(ctx, in, info, handler)
}

// StoreService_ServiceDesc is the grpc.ServiceDesc for StoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.StoreService",
	HandlerType: (*StoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _StoreService_List_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _StoreService_Info_Handler,
		},
		{
			MethodName: "UserList",
			Handler:    _StoreService_UserList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
