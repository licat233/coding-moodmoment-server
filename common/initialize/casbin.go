package initialize

import (
	"moodMoment-server/common/casbinx"
	"moodMoment-server/common/global"
)

func Casbin() {
	global.CASBIN = casbinx.Casbin()
}
