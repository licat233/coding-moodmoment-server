package jwtx

import (
	"moodMoment-server/common/global"
	"moodMoment-server/common/jwtx/model"
)

// var JwtWhitelist model.JwtWhitelistsModel

func NewJwtWhitelist() model.JwtWhitelistsModel {
	return model.NewJwtWhitelistsModel(global.SQLConn, global.CONFIG.Cache.Redis)
}
