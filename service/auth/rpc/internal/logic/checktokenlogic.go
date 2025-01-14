package logic

import (
	"context"

	"togds/service/auth/rpc/authRpc"
	"togds/service/auth/rpc/internal/svc"

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

func (l *CheckTokenLogic) CheckToken(in *authRpc.CheckTokenReq) (*authRpc.CheckTokenResp, error) {
	// todo: add your logic here and delete this line

	return &authRpc.CheckTokenResp{}, nil
}
