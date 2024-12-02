package logic

import (
	"context"

	"docs/welcome/internal/svc"
	"docs/welcome/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WelcomeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWelcomeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WelcomeLogic {
	return &WelcomeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WelcomeLogic) Welcome(req *types.Request) (resp *types.Response, err error) {
	resp = new(types.Response)
	resp.Message = req.Name
	return
}
