package userregisterlogic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"store/app/rpc/user/internal/svc"
	"store/app/rpc/user/pb/user"
	"store/db/dao/model"
	"store/pkg/biz"
	"store/pkg/jwt"
	"store/pkg/util"
	"store/pkg/xcode"
	"time"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (res *user.RegisterRes, err error) {
	var (
		e      error
		userId int64
	)
	code := "200"
	//info := sqls.UsersApi{}
	info := model.UsersApi{}
	token := ""
	password := ""
	res = &user.RegisterRes{
		Result: &user.Response{},
	}
	defer func() {
		res.Result.Code, res.Result.Message = xcode.GetCodeMessage(code)
		if e != nil {
			l.Logger.Errorf("%s 注册用户失败 fail:%s", l.svcCtx.Config.ServiceName, e.Error())
			res.Result.ErrMsg = e.Error()
		} else {
			res.UserId = userId
			res.Token = token
		}
	}()
	info, e = l.svcCtx.MysqlQuery.User.GetFromMobileApi(in.Mobile)
	if e != nil {
		code = xcode.USER_CREAT_FAIL
		goto Result
	}
	if info.UserId > 0 {
		code = xcode.USER_CREAT_MOBILE_FAIL
		goto Result
	}

	password, e = util.HashPassword(in.Password)
	if e != nil {
		code = xcode.USER_CREAT_FAIL
		goto Result
	}
	userId = l.svcCtx.Node.Generate().Int64()
	token, e = jwt.GetJwtToken(in.JwtSecret, time.Now().Unix(), in.Seconds, map[string]interface{}{
		"userId": userId,
	})
	if e != nil {
		code = xcode.USER_CREAT_FAIL
		goto Result
	}
	e = l.svcCtx.MysqlQuery.User.CreatUser(model.Users{
		UserId:   userId,
		Mobile:   in.Mobile,
		Name:     in.Name,
		Password: password,
		Avatar:   "",
	})
	if e != nil {
		code = xcode.USER_CREAT_FAIL
		goto Result
	}
	e = l.svcCtx.CacheConnApi.User.SetInfo(userId, map[string]interface{}{
		"userId": userId,
		"mobile": in.Mobile,
		"name":   in.Name,
		"avatar": "",
	}, in.Seconds)
	if e != nil {
		code = xcode.USER_SET_INFOCACHE_FAIL
		l.Logger.Errorf("%s 存储用户缓存 fail:%s", l.svcCtx.Config.ServiceName, e.Error())
		goto Result
	}
	e = l.svcCtx.BizConn.Set(l.ctx, fmt.Sprintf("%s%s", biz.Biz_Key_USER_TOKEN, token), userId, time.Duration(in.Seconds)*time.Second).Err()
	if e != nil {
		code = xcode.USER_SET_INFOCACHE_FAIL
		l.Logger.Errorf("%s 存储用户token fail:%s", l.svcCtx.Config.ServiceName, e.Error())
	}
Result:
	return res, nil
}
