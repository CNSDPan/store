// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.6
// Source: api.proto

package server

import (
	"context"

	"store/app/rpc/api/internal/logic/apitoken"
	"store/app/rpc/api/internal/svc"
	"store/app/rpc/api/pb/api"
)

type ApiTokenServer struct {
	svcCtx *svc.ServiceContext
	api.UnimplementedApiTokenServer
}

func NewApiTokenServer(svcCtx *svc.ServiceContext) *ApiTokenServer {
	return &ApiTokenServer{
		svcCtx: svcCtx,
	}
}

func (s *ApiTokenServer) CheckAuth(ctx context.Context, in *api.AuthReq) (*api.AuthRes, error) {
	l := apitokenlogic.NewCheckAuthLogic(ctx, s.svcCtx)
	return l.CheckAuth(in)
}
