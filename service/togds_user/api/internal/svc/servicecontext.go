package svc

import (
	"togds/service/togds_user/api/internal/config"
	"togds/service/togds_user/rpc/loginrpcre"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config             config.Config
	TogdsUserRpcClient loginrpcre.LoginRpcre
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:             c,
		TogdsUserRpcClient: loginrpcre.NewLoginRpcre(zrpc.MustNewClient(c.Togdsuserrpc)),
	}
}
