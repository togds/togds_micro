package logic

import (
	"context"
	"fmt"
	"time"

	"togds/service/user/rpc/internal/svc"
	"togds/service/user/rpc/togdsuserrpc"

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

func (l *LoginLogic) Login(in *togdsuserrpc.LoginReq) (*togdsuserrpc.LoginResp, error) {
	// todo: add your logic here and delete this line

	// 生成jwt token
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auths.AccessExpire
	// accessExpire := l.svcCtx.Config.Auth.
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auths.AccessSecret, now, accessExpire, 19191919)
	if err != nil {
		return &togdsuserrpc.LoginResp{}, nil
	}
	err = l.svcCtx.Redis.SetexCtx(l.ctx, fmt.Sprintf("token:%s", jwtToken), in.Username, int(10000))
	fmt.Println(jwtToken)
	return &togdsuserrpc.LoginResp{
		Code: 10000,
		Msg:  "登录成功了",
		Data: jwtToken,
	}, nil
	// return &togdsuserrpc.LoginResp{
	// 	Code: 10000,
	// 	Msg:  "登录成功了",
	// 	Data: "jwtToken",
	// }, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
