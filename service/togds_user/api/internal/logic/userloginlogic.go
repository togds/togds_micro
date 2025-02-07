package logic

import (
	"context"

	"togds/service/togds_user/api/internal/svc"
	"togds/service/togds_user/api/internal/types"
	"togds/service/togds_user/rpc/loginrpcre"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	//TogdsUserRpcClient
	rpcResp, err := l.svcCtx.TogdsUserRpcClient.Login(l.ctx, &loginrpcre.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		Code: rpcResp.Code,
		Msg:  rpcResp.Msg,
		Data: rpcResp.Data,
	}, nil
}
