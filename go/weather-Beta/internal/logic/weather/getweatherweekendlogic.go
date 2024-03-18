package weather

import (
	"context"

	"sun42/weather-Beta/internal/svc"
	"sun42/weather-Beta/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWeatherWeekendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetWeatherWeekendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWeatherWeekendLogic {
	return &GetWeatherWeekendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWeatherWeekendLogic) GetWeatherWeekend(req *types.GetWeatherWeekendReq) (resp *types.GetWeatherWeekendResp, err error) {
	// todo: add your logic here and delete this line

	return
}
