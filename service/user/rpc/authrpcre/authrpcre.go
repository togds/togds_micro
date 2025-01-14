// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.5
// Source: togdsuserrpc.proto

package authrpcre

import (
	"context"

	"togds/service/user/rpc/togdsuserrpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	LoginReq  = togdsuserrpc.LoginReq
	LoginResp = togdsuserrpc.LoginResp

	AuthRpcre interface {
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	}

	defaultAuthRpcre struct {
		cli zrpc.Client
	}
)

func NewAuthRpcre(cli zrpc.Client) AuthRpcre {
	return &defaultAuthRpcre{
		cli: cli,
	}
}

func (m *defaultAuthRpcre) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := togdsuserrpc.NewAuthRpcreClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}
