// Code generated by goctl. DO NOT EDIT.
// Source: store.proto

package server

import (
	"context"

	"store/app/store/rpc/internal/logic/storebecome"
	"store/app/store/rpc/internal/svc"
	"store/app/store/rpc/pb/store"
)

type StoreBecomeServer struct {
	svcCtx *svc.ServiceContext
	store.UnimplementedStoreBecomeServer
}

func NewStoreBecomeServer(svcCtx *svc.ServiceContext) *StoreBecomeServer {
	return &StoreBecomeServer{
		svcCtx: svcCtx,
	}
}

func (s *StoreBecomeServer) CreateStore(ctx context.Context, in *store.CreateStoreReq) (*store.CreateStoreRes, error) {
	l := storebecomelogic.NewCreateStoreLogic(ctx, s.svcCtx)
	return l.CreateStore(in)
}

func (s *StoreBecomeServer) JoinStoreMember(ctx context.Context, in *store.JoinStoreMemberReq) (*store.JoinStoreMemberRes, error) {
	l := storebecomelogic.NewJoinStoreMemberLogic(ctx, s.svcCtx)
	return l.JoinStoreMember(in)
}
