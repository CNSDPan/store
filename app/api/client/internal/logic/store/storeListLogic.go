package store

import (
	"context"

	"store/app/api/client/internal/svc"
	"store/app/api/client/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StoreListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreListLogic {
	return &StoreListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreListLogic) StoreList() (resp *types.StoreListRes, err error) {
	// todo: add your logic here and delete this line

	return
}
