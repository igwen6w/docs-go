package sysUser

import (
	"net/http"

	"docs/app/sys_user_center/cmd/api/internal/logic/sysUser"
	"docs/app/sys_user_center/cmd/api/internal/svc"
	"docs/app/sys_user_center/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// user login
func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sysUser.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
