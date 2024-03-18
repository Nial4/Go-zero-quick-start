package weather

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"sun42/weather-Beta/internal/logic/weather"
	"sun42/weather-Beta/internal/svc"
	"sun42/weather-Beta/internal/types"
)

func GetWeatherDailyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetWeatherDailyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := weather.NewGetWeatherDailyLogic(r.Context(), svcCtx)
		resp, err := l.GetWeatherDaily(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
