package initialize

import (
	"moodMoment-server/common/global"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func SQLConn() {
	global.SQLConn = sqlx.NewMysql(global.CONFIG.Mysql.DataSource())
}
