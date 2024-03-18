package weather

import (
	"context"

	"sun42/weather-Beta/internal/svc"
	"sun42/weather-Beta/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWeatherDailyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetWeatherDailyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWeatherDailyLogic {
	return &GetWeatherDailyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWeatherDailyLogic) GetWeatherDaily(req *types.GetWeatherDailyReq) (resp *types.GetWeatherDailyResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.GetWeatherDailyResp)
	if req.City == "KM" {
		resp.Weather = "晴"
	} else {
		resp.Weather = "雨"
	}

	return
}
