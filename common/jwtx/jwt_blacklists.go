package jwtx

import (
	"moodMoment-server/common/global"
	"moodMoment-server/common/jwtx/model"
)

// var JwtBlacklist model.JwtBlacklistsModel

func NewJwtBlacklist() model.JwtBlacklistsModel {
	return model.NewJwtBlacklistsModel(global.SQLConn, global.CONFIG.Cache.Redis)
}
