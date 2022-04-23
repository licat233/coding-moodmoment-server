package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"moodMoment-server/service/admonitor-api/api/internal/logic"
	"moodMoment-server/service/admonitor-api/api/internal/svc"
)

func GetStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetStatusLogic(r.Context(), svcCtx)
		resp, err := l.GetStatus()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
