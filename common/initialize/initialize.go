package initialize

import (
	"moodMoment-server/common/global"
	"moodMoment-server/common/utils"
	"moodMoment-server/common/zapx"
)

func Initialize() {
	utils.InitTime()
	global.LOG = zapx.NewLogger()
}
