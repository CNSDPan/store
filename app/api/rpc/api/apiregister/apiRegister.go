// Code generated by goctl. DO NOT EDIT.
// Source: api.proto

package apiregister

import (
	"context"

	"store/app/api/rpc/pb/api"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	RegisterReq   = api.RegisterReq
	RegisterRes   = api.RegisterRes
	Response      = api.Response
	StoreInfoReq  = api.StoreInfoReq
	StoreInfoRes  = api.StoreInfoRes
	StoreItem     = api.StoreItem
	StoreListReq  = api.StoreListReq
	StoreListRes  = api.StoreListRes
	StoreUsersReq = api.StoreUsersReq
	StoreUsersRes = api.StoreUsersRes
	StoresMap     = api.StoresMap
	UserInfoReq   = api.UserInfoReq
	UserInfoRes   = api.UserInfoRes
	UserItem      = api.UserItem
	UserLoginReq  = api.UserLoginReq
	UserLoginRes  = api.UserLoginRes
	UsersMap      = api.UsersMap

	ApiRegister interface {
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRes, error)
	}

	defaultApiRegister struct {
		cli zrpc.Client
	}
)

func NewApiRegister(cli zrpc.Client) ApiRegister {
	return &defaultApiRegister{
		cli: cli,
	}
}

func (m *defaultApiRegister) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRes, error) {
	client := api.NewApiRegisterClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}