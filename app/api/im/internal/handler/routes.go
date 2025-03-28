// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.6

package handler

import (
	"net/http"
	"time"

	ws "store/app/api/im/internal/handler/ws"
	"store/app/api/im/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.XHeaderMiddleware, serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					// socket 连接
					Method:  http.MethodGet,
					Path:    "/socket",
					Handler: ws.SocketHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/v1/ws"),
		rest.WithTimeout(30000*time.Millisecond),
	)
}
