// Code generated by goctl. DO NOT EDIT.
// Source: store.proto

package server

import (
	"context"

	"store/app/store/rpc/internal/logic/storeservice"
	"store/app/store/rpc/internal/svc"
	"store/app/store/rpc/pb/store"
)

type StoreServiceServer struct {
	svcCtx *svc.ServiceContext
	store.UnimplementedStoreServiceServer
}

func NewStoreServiceServer(svcCtx *svc.ServiceContext) *StoreServiceServer {
	return &StoreServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *StoreServiceServer) CreateStore(ctx context.Context, in *store.CreateStoreReq) (*store.CreateStoreRes, error) {
	l := storeservicelogic.NewCreateStoreLogic(ctx, s.svcCtx)
	return l.CreateStore(in)
}

func (s *StoreServiceServer) JoinStoreMember(ctx context.Context, in *store.JoinStoreMemberReq) (*store.JoinStoreMemberRes, error) {
	l := storeservicelogic.NewJoinStoreMemberLogic(ctx, s.svcCtx)
	return l.JoinStoreMember(in)
}