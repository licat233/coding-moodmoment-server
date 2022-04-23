package initialize

import (
	"moodMoment-server/common/global"
	"moodMoment-server/common/gormx"
)

func GormDB() {
	global.GORM = gormx.CreateGorm()
}
