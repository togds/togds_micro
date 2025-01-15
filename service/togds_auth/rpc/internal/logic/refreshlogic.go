package logic

import (
	"context"

	"togds/service/togds_auth/rpc/internal/svc"
	"togds/service/togds_auth/rpc/togds_auth"

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

func (l *RefreshLogic) Refresh(in *togds_auth.RefreshReq) (*togds_auth.RefreshResp, error) {
	// todo: add your logic here and delete this line

	return &togds_auth.RefreshResp{}, nil
}
