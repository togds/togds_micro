package svc

import (
	"togds/service/user/api/internal/config"
	"togds/service/user/rpc/authrpcre"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config             config.Config
	Togdsuserrpcclient authrpcre.AuthRpcre
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:             c,
		Togdsuserrpcclient: authrpcre.NewAuthRpcre(zrpc.MustNewClient(c.Togdsuserrpc)),
	}
}
