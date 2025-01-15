package logic

import (
	"context"

	"togds/service/togds_auth/rpc/internal/svc"
	"togds/service/togds_auth/rpc/togds_auth"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckTokenLogic {
	return &CheckTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckTokenLogic) CheckToken(in *togds_auth.CheckTokenReq) (*togds_auth.CheckTokenResp, error) {
	// todo: add your logic here and delete this line

	return &togds_auth.CheckTokenResp{}, nil
}
