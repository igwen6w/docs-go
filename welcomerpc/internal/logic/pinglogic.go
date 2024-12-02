package logic

import (
	"context"

	"docs/welcomerpc/internal/svc"
	"docs/welcomerpc/welcomerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *welcomerpc.Request) (*welcomerpc.Response, error) {
	return &welcomerpc.Response{
		Pong: "pong",
	}, nil
}
