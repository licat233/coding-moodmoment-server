package initialize

import (
	"moodMoment-server/common/global"
	"moodMoment-server/common/redisx"
)

func Redis() {
	global.REDIS = redisx.CreateRedis()
}

func RedisX() {
	global.RedisX = redisx.CreateRedisX()
}
