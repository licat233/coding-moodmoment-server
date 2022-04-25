package svc

import (
	"moodMoment-server/service/admonitor-api/api/internal/config"
	"moodMoment-server/service/admonitor-api/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	PopinModel model.PopinModel
	LineModel  model.LineModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		PopinModel: model.NewPopinModel(conn),
		LineModel:  model.NewLineModel(conn),
	}
}
