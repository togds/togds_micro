package svc

import (
	"togds/service/togds_user/rpc/internal/config"
	"togds/service/togds_user/rpc/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Redis
	Model  model.TUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Redis:  redis.MustNewRedis(c.RedisConfig),
		Model:  model.NewTUserModel(sqlx.NewMysql(c.DataSource)),
	}
}
