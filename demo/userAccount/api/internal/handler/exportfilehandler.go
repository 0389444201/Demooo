package handler

import (
	"net/http"

	"demo/userAccount/api/internal/logic"
	"demo/userAccount/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ExportFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewExportFileLogic(r.Context(), svcCtx)
		resp, err := l.ExportFile()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
