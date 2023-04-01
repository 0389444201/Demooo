package handler

import (
	"net/http"

	"demo/userAccount/api/internal/logic"
	"demo/userAccount/api/internal/svc"
	"demo/userAccount/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AccountRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAccountRegisterLogic(r.Context(), svcCtx)
		resp, err := l.AccountRegister(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
