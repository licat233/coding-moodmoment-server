package logic

import (
	"context"

	"moodMoment-server/service/admonitor-api/api/internal/svc"
	"moodMoment-server/service/admonitor-api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStatusLogic {
	return &GetStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStatusLogic) GetStatus() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
