package main

import (
	"flag"
	"fmt"

	"togds/service/togds_user/rpc/internal/config"
	"togds/service/togds_user/rpc/internal/server"
	"togds/service/togds_user/rpc/internal/svc"
	"togds/service/togds_user/rpc/togds_userRpc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/togdsuserrpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		togds_userRpc.RegisterLoginRpcreServer(grpcServer, server.NewLoginRpcreServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
