package logic

import (
	"context"

	"togds/service/auth/rpc/authRpc"
	"togds/service/auth/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshLogic {
	return &RefreshLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefreshLogic) Refresh(in *authRpc.RefreshReq) (*authRpc.RefreshResp, error) {
	// todo: add your logic here and delete this line

	return &authRpc.RefreshResp{}, nil
}
