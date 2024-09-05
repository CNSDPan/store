// Code generated by goctl. DO NOT EDIT.
// Source: api.proto

package storeservice

import (
	"context"

	"store/app/api/rpc/pb/api"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
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

	StoreService interface {
		List(ctx context.Context, in *StoreListReq, opts ...grpc.CallOption) (*StoreListRes, error)
		Info(ctx context.Context, in *StoreInfoReq, opts ...grpc.CallOption) (*StoreInfoRes, error)
		UserList(ctx context.Context, in *StoreUsersReq, opts ...grpc.CallOption) (*StoreUsersRes, error)
	}

	defaultStoreService struct {
		cli zrpc.Client
	}
)

func NewStoreService(cli zrpc.Client) StoreService {
	return &defaultStoreService{
		cli: cli,
	}
}

func (m *defaultStoreService) List(ctx context.Context, in *StoreListReq, opts ...grpc.CallOption) (*StoreListRes, error) {
	client := api.NewStoreServiceClient(m.cli.Conn())
	return client.List(ctx, in, opts...)
}

func (m *defaultStoreService) Info(ctx context.Context, in *StoreInfoReq, opts ...grpc.CallOption) (*StoreInfoRes, error) {
	client := api.NewStoreServiceClient(m.cli.Conn())
	return client.Info(ctx, in, opts...)
}

func (m *defaultStoreService) UserList(ctx context.Context, in *StoreUsersReq, opts ...grpc.CallOption) (*StoreUsersRes, error) {
	client := api.NewStoreServiceClient(m.cli.Conn())
	return client.UserList(ctx, in, opts...)
}