package storebecomelogic

import (
	"context"
	"store/app/chat/rpc/chat/socket"
	"store/pkg/consts"
	"store/pkg/xcode"

	"store/app/store/rpc/internal/svc"
	"store/app/store/rpc/pb/store"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinStoreMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJoinStoreMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinStoreMemberLogic {
	return &JoinStoreMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JoinStoreMemberLogic) JoinStoreMember(in *store.JoinStoreMemberReq) (res *store.JoinStoreMemberRes, err error) {
	var (
		e    error
		code = "200"
		has  int64
	)
	res = &store.JoinStoreMemberRes{
		Result: &store.Response{},
	}
	defer func() {
		res.Result.Code, res.Result.Message = xcode.GetCodeMessage(code)
		if e != nil {
			l.Logger.Errorf("%s 加入门店会员 fail:%s", l.svcCtx.Config.ServiceName, e.Error())
			res.Result.ErrMsg = e.Error()
		}
	}()
	if l.svcCtx.StoreMemberModel.GetMemberContacts(in.StoreId) >= 50 {
		code = xcode.STORE_MEMBER_JOIN_FULL
		return res, err
	}
	has, e = l.svcCtx.StoreMemberModel.MemberJoin(in.StoreId, in.UserId, l.svcCtx.Node.Generate().Int64())
	if has > 0 {
		code = xcode.STORE_MEMBER_JOINED
		return res, err
	} else if e != nil {
		code = xcode.STORE_MEMBER_JOIN_FAIL
		return res, err
	}
	userMap, chatErr := l.svcCtx.CacheConnApi.User.GetInfo(in.UserId)
	if chatErr != nil {
		l.Logger.Errorf("%s 获取用户信息失败 fail:%s", l.svcCtx.Config.ServiceName, chatErr.Error())
		return res, nil
	}

	_, chatErr = l.svcCtx.ChatRpcCl.Socket.BroadcastBecomeMsg(context.Background(), &socket.BroadcastReq{
		Operate:       int32(consts.OperatePublic),
		Method:        consts.MethodBecome,
		StoreId:       in.StoreId,
		SendUserId:    in.UserId,
		SendUserName:  userMap["name"],
		ReceiveUserId: 0,
		Extend:        "",
		Body:          "",
	})
	if chatErr != nil {
		l.Logger.Errorf("%s 广播成为会员消息 fail:%s", l.svcCtx.Config.ServiceName, chatErr.Error())
	}
	return res, nil
}
