// Code generated by goctl. DO NOT EDIT.
// Source: store.proto

package storebecome

import (
	"context"

	"store/app/store/rpc/pb/store"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateStoreReq     = store.CreateStoreReq
	CreateStoreRes     = store.CreateStoreRes
	JoinStoreMemberReq = store.JoinStoreMemberReq
	JoinStoreMemberRes = store.JoinStoreMemberRes
	Response           = store.Response

	StoreBecome interface {
		CreateStore(ctx context.Context, in *CreateStoreReq, opts ...grpc.CallOption) (*CreateStoreRes, error)
		JoinStoreMember(ctx context.Context, in *JoinStoreMemberReq, opts ...grpc.CallOption) (*JoinStoreMemberRes, error)
	}

	defaultStoreBecome struct {
		cli zrpc.Client
	}
)

func NewStoreBecome(cli zrpc.Client) StoreBecome {
	return &defaultStoreBecome{
		cli: cli,
	}
}

func (m *defaultStoreBecome) CreateStore(ctx context.Context, in *CreateStoreReq, opts ...grpc.CallOption) (*CreateStoreRes, error) {
	client := store.NewStoreBecomeClient(m.cli.Conn())
	return client.CreateStore(ctx, in, opts...)
}

func (m *defaultStoreBecome) JoinStoreMember(ctx context.Context, in *JoinStoreMemberReq, opts ...grpc.CallOption) (*JoinStoreMemberRes, error) {
	client := store.NewStoreBecomeClient(m.cli.Conn())
	return client.JoinStoreMember(ctx, in, opts...)
}