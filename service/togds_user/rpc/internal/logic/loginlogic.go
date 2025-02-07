package logic

import (
	"context"
	"fmt"
	"strconv"
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
	userinfo, err := l.svcCtx.Model.FindOneUsername(l.ctx, in.Username)
	if err != nil || len(in.Password) != 32 || userinfo.Password != in.Password {
		return &togds_userRpc.LoginResp{
			Code: 10004,
			Msg:  "登录失败",
			Data: "err",
		}, nil
	}

	// 查询token是否存在
	redisToken, err := l.svcCtx.Redis.GetCtx(l.ctx, strconv.Itoa(int(userinfo.UserId)))
	if err != nil {
		return &togds_userRpc.LoginResp{
			Code: 10004,
			Msg:  "登录失败11111",
			Data: "err",
		}, nil
	}

	if redisToken != "" {
		return &togds_userRpc.LoginResp{
			Code: 10000,
			Msg:  "登录成功",
			Data: redisToken,
		}, nil
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auths.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auths.AccessSecret, now, accessExpire, 19191919)
	if err != nil {
		return &togds_userRpc.LoginResp{
			Code: 10003,
			Msg:  "生成token失败",
			Data: "err",
		}, nil
	}
	err = l.svcCtx.Redis.SetCtx(l.ctx, strconv.Itoa(int(userinfo.UserId)), jwtToken)
	if err != nil {
		return &togds_userRpc.LoginResp{
			Code: 10001,
			Msg:  "插入redis失败",
			Data: jwtToken,
		}, nil
	}
	err = l.svcCtx.Redis.SetCtx(l.ctx, jwtToken, strconv.Itoa(int(userinfo.UserId)))
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
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
