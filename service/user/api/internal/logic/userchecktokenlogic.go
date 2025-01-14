package logic

import (
	"context"

	"togds/service/user/api/internal/svc"
	"togds/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCheckTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCheckTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCheckTokenLogic {
	return &UserCheckTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCheckTokenLogic) UserCheckToken(req *types.UserAddReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line

	return
}
