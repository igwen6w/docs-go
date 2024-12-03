package sysUser

import (
	"context"

	"docs/app/sys_user_center/cmd/api/internal/svc"
	"docs/app/sys_user_center/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// user login
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	m := map[string]string{
		"admin":   "xxxxxx-xx-xx",
		"igwen6w": "aaaaaa-aa-aa",
	}

	accessToken := "unknown"
	if token, ok := m[req.Account]; ok {
		accessToken = token
	}

	return &types.LoginResp{
		AccessToken: accessToken,
	}, nil
}
