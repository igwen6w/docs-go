package handler

import (
	"net/http"

	"docs/welcome/internal/logic"
	"docs/welcome/internal/svc"
	"docs/welcome/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func WelcomeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewWelcomeLogic(r.Context(), svcCtx)
		resp, err := l.Welcome(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
