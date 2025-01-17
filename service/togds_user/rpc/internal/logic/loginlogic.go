package logic

import (
	"context"
	"fmt"
	"time"

	"togds/service/togds_user/rpc/internal/svc"
	"togds/service/togds_user/rpc/togds_userRpc"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *togds_userRpc.LoginReq) (*togds_userRpc.LoginResp, error) {
	// 查询用户
	_, err := l.svcCtx.Model.FindOneUsername(l.ctx, in.Username)
	// user, err := l.svcCtx.Model.FindOne(l.ctx, 3)
	if err != nil {
		return &togds_userRpc.LoginResp{
			Code: 10004,
			Msg:  "用户不存在",
			Data: err.Error(),
		}, nil
	}
	// todo: add your logic here and delete this line
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auths.AccessExpire
	// accessExpire := l.svcCtx.Config.Auth.
	fmt.Println("开始生成token")
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auths.AccessSecret, now, accessExpire, 19191919)
	if err != nil {
		return &togds_userRpc.LoginResp{
			Code: 10003,
			Msg:  "生成token失败",
			Data: "err",
		}, nil
	}
	err = l.svcCtx.Redis.SetexCtx(l.ctx, fmt.Sprintf("token:%s", jwtToken), in.Username, int(10000))
	if err != nil {
		return &togds_userRpc.LoginResp{
			Code: 10001,
			Msg:  "插入redis失败",
			Data: jwtToken,
		}, nil
	}
	fmt.Println(jwtToken)
	return &togds_userRpc.LoginResp{
		Code: 10000,
		Msg:  "登录成功了",
		Data: jwtToken,
	}, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	fmt.Println("开始生成tokensssss")
	token := jwt.New(jwt.SigningMethodHS256)
	fmt.Println("开始生成tokensssssddddddddddddd", token)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
