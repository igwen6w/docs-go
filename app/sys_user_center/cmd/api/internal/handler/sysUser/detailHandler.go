package sysUser

import (
	"net/http"

	"docs/app/sys_user_center/cmd/api/internal/logic/sysUser"
	"docs/app/sys_user_center/cmd/api/internal/svc"
	"docs/app/sys_user_center/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// get user info
func DetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sysUser.NewDetailLogic(r.Context(), svcCtx)
		resp, err := l.Detail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
