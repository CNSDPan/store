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
	ApiUser_Login_FullMethodName = "/api.ApiUser/Login"
	ApiUser_Info_FullMethodName  = "/api.ApiUser/Info"
)

// ApiUserClient is the client API for ApiUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiUserClient interface {
	Login(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginRes, error)
	Info(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoRes, error)
}

type apiUserClient struct {
	cc grpc.ClientConnInterface
}

func NewApiUserClient(cc grpc.ClientConnInterface) ApiUserClient {
	return &apiUserClient{cc}
}

func (c *apiUserClient) Login(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginRes, error) {
	out := new(UserLoginRes)
	err := c.cc.Invoke(ctx, ApiUser_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiUserClient) Info(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoRes, error) {
	out := new(UserInfoRes)
	err := c.cc.Invoke(ctx, ApiUser_Info_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiUserServer is the server API for ApiUser service.
// All implementations must embed UnimplementedApiUserServer
// for forward compatibility
type ApiUserServer interface {
	Login(context.Context, *UserLoginReq) (*UserLoginRes, error)
	Info(context.Context, *UserInfoReq) (*UserInfoRes, error)
	mustEmbedUnimplementedApiUserServer()
}

// UnimplementedApiUserServer must be embedded to have forward compatible implementations.
type UnimplementedApiUserServer struct {
}

func (UnimplementedApiUserServer) Login(context.Context, *UserLoginReq) (*UserLoginRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedApiUserServer) Info(context.Context, *UserInfoReq) (*UserInfoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedApiUserServer) mustEmbedUnimplementedApiUserServer() {}

// UnsafeApiUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiUserServer will
// result in compilation errors.
type UnsafeApiUserServer interface {
	mustEmbedUnimplementedApiUserServer()
}

func RegisterApiUserServer(s grpc.ServiceRegistrar, srv ApiUserServer) {
	s.RegisterService(&ApiUser_ServiceDesc, srv)
}

func _ApiUser_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiUserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiUser_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiUserServer).Login(ctx, req.(*UserLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiUser_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiUserServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiUser_Info_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiUserServer).Info(ctx, req.(*UserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiUser_ServiceDesc is the grpc.ServiceDesc for ApiUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ApiUser",
	HandlerType: (*ApiUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _ApiUser_Login_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _ApiUser_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

const (
	ApiStore_List_FullMethodName           = "/api.ApiStore/List"
	ApiStore_Info_FullMethodName           = "/api.ApiStore/Info"
	ApiStore_MemberUserList_FullMethodName = "/api.ApiStore/MemberUserList"
	ApiStore_MyAllStore_FullMethodName     = "/api.ApiStore/MyAllStore"
)

// ApiStoreClient is the client API for ApiStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiStoreClient interface {
	List(ctx context.Context, in *StoreListReq, opts ...grpc.CallOption) (*StoreListRes, error)
	Info(ctx context.Context, in *StoreInfoReq, opts ...grpc.CallOption) (*StoreInfoRes, error)
	MemberUserList(ctx context.Context, in *MemberUsersItemReq, opts ...grpc.CallOption) (*MemberUsersItemRes, error)
	MyAllStore(ctx context.Context, in *MyAllStoreIdReq, opts ...grpc.CallOption) (*MyAllStoreIdRes, error)
}

type apiStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewApiStoreClient(cc grpc.ClientConnInterface) ApiStoreClient {
	return &apiStoreClient{cc}
}

func (c *apiStoreClient) List(ctx context.Context, in *StoreListReq, opts ...grpc.CallOption) (*StoreListRes, error) {
	out := new(StoreListRes)
	err := c.cc.Invoke(ctx, ApiStore_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiStoreClient) Info(ctx context.Context, in *StoreInfoReq, opts ...grpc.CallOption) (*StoreInfoRes, error) {
	out := new(StoreInfoRes)
	err := c.cc.Invoke(ctx, ApiStore_Info_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiStoreClient) MemberUserList(ctx context.Context, in *MemberUsersItemReq, opts ...grpc.CallOption) (*MemberUsersItemRes, error) {
	out := new(MemberUsersItemRes)
	err := c.cc.Invoke(ctx, ApiStore_MemberUserList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiStoreClient) MyAllStore(ctx context.Context, in *MyAllStoreIdReq, opts ...grpc.CallOption) (*MyAllStoreIdRes, error) {
	out := new(MyAllStoreIdRes)
	err := c.cc.Invoke(ctx, ApiStore_MyAllStore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiStoreServer is the server API for ApiStore service.
// All implementations must embed UnimplementedApiStoreServer
// for forward compatibility
type ApiStoreServer interface {
	List(context.Context, *StoreListReq) (*StoreListRes, error)
	Info(context.Context, *StoreInfoReq) (*StoreInfoRes, error)
	MemberUserList(context.Context, *MemberUsersItemReq) (*MemberUsersItemRes, error)
	MyAllStore(context.Context, *MyAllStoreIdReq) (*MyAllStoreIdRes, error)
	mustEmbedUnimplementedApiStoreServer()
}

// UnimplementedApiStoreServer must be embedded to have forward compatible implementations.
type UnimplementedApiStoreServer struct {
}

func (UnimplementedApiStoreServer) List(context.Context, *StoreListReq) (*StoreListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedApiStoreServer) Info(context.Context, *StoreInfoReq) (*StoreInfoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedApiStoreServer) MemberUserList(context.Context, *MemberUsersItemReq) (*MemberUsersItemRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberUserList not implemented")
}
func (UnimplementedApiStoreServer) MyAllStore(context.Context, *MyAllStoreIdReq) (*MyAllStoreIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MyAllStore not implemented")
}
func (UnimplementedApiStoreServer) mustEmbedUnimplementedApiStoreServer() {}

// UnsafeApiStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiStoreServer will
// result in compilation errors.
type UnsafeApiStoreServer interface {
	mustEmbedUnimplementedApiStoreServer()
}

func RegisterApiStoreServer(s grpc.ServiceRegistrar, srv ApiStoreServer) {
	s.RegisterService(&ApiStore_ServiceDesc, srv)
}

func _ApiStore_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiStoreServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiStore_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiStoreServer).List(ctx, req.(*StoreListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiStore_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiStoreServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiStore_Info_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiStoreServer).Info(ctx, req.(*StoreInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiStore_MemberUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberUsersItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiStoreServer).MemberUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiStore_MemberUserList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiStoreServer).MemberUserList(ctx, req.(*MemberUsersItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiStore_MyAllStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MyAllStoreIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiStoreServer).MyAllStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiStore_MyAllStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiStoreServer).MyAllStore(ctx, req.(*MyAllStoreIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiStore_ServiceDesc is the grpc.ServiceDesc for ApiStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ApiStore",
	HandlerType: (*ApiStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _ApiStore_List_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _ApiStore_Info_Handler,
		},
		{
			MethodName: "MemberUserList",
			Handler:    _ApiStore_MemberUserList_Handler,
		},
		{
			MethodName: "MyAllStore",
			Handler:    _ApiStore_MyAllStore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

const (
	ApiToken_CheckAuth_FullMethodName = "/api.ApiToken/CheckAuth"
)

// ApiTokenClient is the client API for ApiToken service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiTokenClient interface {
	CheckAuth(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthRes, error)
}

type apiTokenClient struct {
	cc grpc.ClientConnInterface
}

func NewApiTokenClient(cc grpc.ClientConnInterface) ApiTokenClient {
	return &apiTokenClient{cc}
}

func (c *apiTokenClient) CheckAuth(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthRes, error) {
	out := new(AuthRes)
	err := c.cc.Invoke(ctx, ApiToken_CheckAuth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiTokenServer is the server API for ApiToken service.
// All implementations must embed UnimplementedApiTokenServer
// for forward compatibility
type ApiTokenServer interface {
	CheckAuth(context.Context, *AuthReq) (*AuthRes, error)
	mustEmbedUnimplementedApiTokenServer()
}

// UnimplementedApiTokenServer must be embedded to have forward compatible implementations.
type UnimplementedApiTokenServer struct {
}

func (UnimplementedApiTokenServer) CheckAuth(context.Context, *AuthReq) (*AuthRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAuth not implemented")
}
func (UnimplementedApiTokenServer) mustEmbedUnimplementedApiTokenServer() {}

// UnsafeApiTokenServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiTokenServer will
// result in compilation errors.
type UnsafeApiTokenServer interface {
	mustEmbedUnimplementedApiTokenServer()
}

func RegisterApiTokenServer(s grpc.ServiceRegistrar, srv ApiTokenServer) {
	s.RegisterService(&ApiToken_ServiceDesc, srv)
}

func _ApiToken_CheckAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiTokenServer).CheckAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiToken_CheckAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiTokenServer).CheckAuth(ctx, req.(*AuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiToken_ServiceDesc is the grpc.ServiceDesc for ApiToken service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiToken_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ApiToken",
	HandlerType: (*ApiTokenServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckAuth",
			Handler:    _ApiToken_CheckAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
