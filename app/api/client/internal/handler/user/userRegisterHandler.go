package user

import (
	"net/http"
	"store/pkg/xcode"

	"github.com/zeromicro/go-zero/rest/httpx"
	"store/app/api/client/internal/logic/user"
	"store/app/api/client/internal/svc"
	"store/app/api/client/internal/types"
)

func UserRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			code, msg := xcode.GetCodeMessage(xcode.RESPONSE_NOT_FOUND)
			httpx.OkJsonCtx(r.Context(), w, types.JSONResponseCtx{
				ErrMsg:  err.Error(),
				Code:    code,
				Message: msg,
				Data:    map[string]interface{}{},
			})
			return
		}

		l := user.NewUserRegisterLogic(r.Context(), svcCtx)
		res, resp, err := l.UserRegister(&req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.JSONResponseCtx{
				ErrMsg:  res.ErrMsg,
				Code:    res.Code,
				Message: res.Message,
				Data:    map[string]interface{}{},
			})
		} else {
			httpx.OkJsonCtx(r.Context(), w, types.JSONResponseCtx{
				ErrMsg:  res.ErrMsg,
				Code:    res.Code,
				Message: res.Message,
				Data:    resp,
			})
		}
	}
}
