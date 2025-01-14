package logic

import (
	"context"

	"togds/service/user/api/internal/svc"
	"togds/service/user/api/internal/types"
	"togds/service/user/rpc/authrpcre"

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
	rpcResp, err := l.svcCtx.Togdsuserrpcclient.Login(l.ctx, &authrpcre.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		Code: 10000,
		Msg:  "登录成功",
		Data: rpcResp.Data,
	}, nil
}
