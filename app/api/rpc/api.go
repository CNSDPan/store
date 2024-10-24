package main

import (
	"flag"
	"fmt"

	"store/app/api/rpc/internal/config"
	apistoreServer "store/app/api/rpc/internal/server/apistore"
	apitokenServer "store/app/api/rpc/internal/server/apitoken"
	apiuserServer "store/app/api/rpc/internal/server/apiuser"
	"store/app/api/rpc/internal/svc"
	"store/app/api/rpc/pb/api"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		api.RegisterApiUserServer(grpcServer, apiuserServer.NewApiUserServer(ctx))
		api.RegisterApiStoreServer(grpcServer, apistoreServer.NewApiStoreServer(ctx))
		api.RegisterApiTokenServer(grpcServer, apitokenServer.NewApiTokenServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
