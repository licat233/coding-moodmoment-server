package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"moodMoment-server/service/admonitor-api/api/internal/logic"
	"moodMoment-server/service/admonitor-api/api/internal/svc"
	"moodMoment-server/service/admonitor-api/api/internal/types"
)

func GetStatisticsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetStatisticsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetStatisticsLogic(r.Context(), svcCtx)
		resp, err := l.GetStatistics(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
