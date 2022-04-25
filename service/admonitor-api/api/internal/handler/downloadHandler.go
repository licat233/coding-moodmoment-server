package handler

import (
	"net/http"

	"moodMoment-server/service/admonitor-api/api/internal/logic"
	"moodMoment-server/service/admonitor-api/api/internal/svc"
	"moodMoment-server/service/admonitor-api/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DownloadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDownloadLogic(r.Context(), svcCtx)
		body, err := l.Download(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			w.Write(body)
			// httpx.Ok(w)
		}
	}
}
