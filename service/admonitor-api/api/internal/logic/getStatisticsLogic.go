package logic

import (
	"context"
	"errors"
	"strings"
	"time"

	"moodMoment-server/common/errorx"
	"moodMoment-server/common/respx"
	"moodMoment-server/common/utils"
	"moodMoment-server/service/admonitor-api/api/internal/svc"
	"moodMoment-server/service/admonitor-api/api/internal/types"
	"moodMoment-server/service/admonitor-api/model"

	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type GetStatisticsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStatisticsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStatisticsLogic {
	return &GetStatisticsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStatisticsLogic) GetStatistics(req *types.GetStatisticsRequest) (any, error) {
	channel := req.Channel
	startDate, err := utils.ParseDate(req.StartDate)
	if err != nil {
		return nil, errors.New("'startDate' param error")
	}
	endDate, err := utils.ParseDate(req.EndDate)
	if err != nil {
		return nil, errors.New("'endDate' param error")
	}
	if startDate.After(endDate) {
		return nil, errors.New("开始时间大于结束时间")
	}
	now := time.Now().Local()
	if startDate.After(now) {
		return respx.SingleResp("无此时间段的数据", nil), nil
	}
	account := req.Account
	if strings.TrimSpace(account) == "" {
		return nil, errors.New("'account' param error")
	}
	switch channel {
	case "line":
		records, err := l.svcCtx.LineModel.FindsByDate(l.ctx, startDate, endDate, account)
		return lineRecordsHandler(records, err)
	case "popin":
		records, err := l.svcCtx.PopinModel.FindsByDate(l.ctx, startDate, endDate, account)
		return popinRecordsHandler(records, err)
	}
	return nil, errors.New("'channel' param error")
}

func lineRecordsHandler(records []*model.Line, err error) (any, error) {
	if errors.Is(err, sqlx.ErrNotFound) {
		return respx.SingleResp("无此时间段的数据", nil), nil
	}
	if err != nil {
		return nil, errorx.InternalError("服务器错误，请稍后再试", err)
	}
	line := &model.Line{}
	for _, record := range records {
		calculateLineData(line, record)
	}
	return respx.SingleResp("请求成功", line), nil
}

func popinRecordsHandler(records []*model.Popin, err error) (any, error) {
	if errors.Is(err, sqlx.ErrNotFound) {
		return respx.SingleResp("无此时间段的数据", nil), nil
	}
	if err != nil {
		return nil, errorx.InternalError("服务器错误，请稍后再试", err)
	}
	popin := &model.Popin{}
	for _, record := range records {
		calculatePopinData(popin, record)
	}
	return respx.SingleResp("请求成功", popin), nil
}

func calculateLineData(target, source *model.Line) {
	cost, _ := decimal.NewFromFloat(target.Cost).Add(decimal.NewFromFloat(source.Cost)).Round(2).Float64()
	imp := target.Imp + source.Imp
	click := target.Click + source.Click
	cv := target.Cv + source.Cv
	ctr := utils.CtrCalculation(click, imp)
	cvr := utils.CvrCalculation(cv, click)
	cpa := utils.CpaCalculation(cost, cv)
	target.Cost = cost
	target.Imp = imp
	target.Click = click
	target.Cv = cv
	target.Ctr = ctr
	target.Cvr = cvr
	target.Cpa = cpa
	target.Tp = source.Tp
	target.Currency = source.Currency
}

func calculatePopinData(target, source *model.Popin) {
	cost, _ := decimal.NewFromFloat(target.Cost).Add(decimal.NewFromFloat(source.Cost)).Round(2).Float64()
	imp := target.Imp + source.Imp
	click := target.Click + source.Click
	cv := target.Cv + source.Cv
	ctr := utils.CtrCalculation(click, imp)
	cvr := utils.CvrCalculation(cv, click)
	cpa := utils.CpaCalculation(cost, cv)
	target.Cost = cost
	target.Imp = imp
	target.Click = click
	target.Cv = cv
	target.Ctr = ctr
	target.Cvr = cvr
	target.Cpa = cpa
	target.Tp = source.Tp
	target.Currency = source.Currency
}
