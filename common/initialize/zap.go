package initialize

import (
	"moodMoment-server/common/global"
	"moodMoment-server/common/zapx"
)

func Zap() {
	global.ZAP = zapx.Zap()
}
