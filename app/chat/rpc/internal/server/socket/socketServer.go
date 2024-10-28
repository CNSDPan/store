// Code generated by goctl. DO NOT EDIT.
// Source: chat.proto

package server

import (
	"context"

	"store/app/chat/rpc/internal/logic/socket"
	"store/app/chat/rpc/internal/svc"
	"store/app/chat/rpc/pb/chat"
)

type SocketServer struct {
	svcCtx *svc.ServiceContext
	chat.UnimplementedSocketServer
}

func NewSocketServer(svcCtx *svc.ServiceContext) *SocketServer {
	return &SocketServer{
		svcCtx: svcCtx,
	}
}

func (s *SocketServer) BroadcastMsg(ctx context.Context, in *chat.BroadcastReq) (*chat.Response, error) {
	l := socketlogic.NewBroadcastMsgLogic(ctx, s.svcCtx)
	return l.BroadcastMsg(in)
}