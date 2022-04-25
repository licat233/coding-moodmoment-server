package logic

import (
	"context"
	"errors"
	"io/ioutil"
	"path"
	"strings"

	"moodMoment-server/common/utils"
	"moodMoment-server/service/admonitor-api/api/internal/svc"
	"moodMoment-server/service/admonitor-api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadLogic {
	return &DownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadLogic) Download(req *types.DownloadReq) ([]byte, error) {
	filename := req.File
	if filename = strings.TrimSpace(filename); filename == "" {
		return nil, errors.New("empty file name")
	}
	staticPath := path.Join(utils.CurrentProjectRootPath(), "/service/admonitor-api/static/", filename)
	body, err := ioutil.ReadFile(staticPath)
	if err != nil {
		return nil, err
	}
	return body, nil
}
