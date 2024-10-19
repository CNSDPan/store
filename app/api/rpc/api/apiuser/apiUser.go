// Code generated by goctl. DO NOT EDIT.
// Source: api.proto

package apiuser

import (
	"context"

	"store/app/api/rpc/pb/api"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AuthReq            = api.AuthReq
	AuthRes            = api.AuthRes
	MemberUsersItemReq = api.MemberUsersItemReq
	MemberUsersItemRes = api.MemberUsersItemRes
	MyAllStoreIdReq    = api.MyAllStoreIdReq
	MyAllStoreIdRes    = api.MyAllStoreIdRes
	Response           = api.Response
	StoreInfoReq       = api.StoreInfoReq
	StoreInfoRes       = api.StoreInfoRes
	StoreItem          = api.StoreItem
	StoreListReq       = api.StoreListReq
	StoreListRes       = api.StoreListRes
	StoresMap          = api.StoresMap
	UserInfoReq        = api.UserInfoReq
	UserInfoRes        = api.UserInfoRes
	UserItem           = api.UserItem
	UserLoginReq       = api.UserLoginReq
	UserLoginRes       = api.UserLoginRes
	UsersMap           = api.UsersMap

	ApiUser interface {
		Login(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginRes, error)
		Info(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoRes, error)
	}

	defaultApiUser struct {
		cli zrpc.Client
	}
)

func NewApiUser(cli zrpc.Client) ApiUser {
	return &defaultApiUser{
		cli: cli,
	}
}

func (m *defaultApiUser) Login(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginRes, error) {
	client := api.NewApiUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultApiUser) Info(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoRes, error) {
	client := api.NewApiUserClient(m.cli.Conn())
	return client.Info(ctx, in, opts...)
}
