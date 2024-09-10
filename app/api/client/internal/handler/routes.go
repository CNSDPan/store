// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	"time"

	store "store/app/api/client/internal/handler/store"
	user "store/app/api/client/internal/handler/user"
	"store/app/api/client/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.XHeaderMiddleware, serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/store/info",
					Handler: store.StoreInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/store/list",
					Handler: store.StoreListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/store/user/list",
					Handler: store.StoreUserListHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/v1/client"),
		rest.WithTimeout(3000*time.Millisecond),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.XHeaderMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/login",
					Handler: user.UserLoginHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/register",
					Handler: user.UserRegisterHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/v1/client"),
		rest.WithTimeout(3000*time.Millisecond),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.XHeaderMiddleware, serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/info",
					Handler: user.UserInfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/v1/client"),
		rest.WithTimeout(3000*time.Millisecond),
	)
}
