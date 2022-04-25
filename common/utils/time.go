package utils

import "time"

var LOC, _ = time.LoadLocation("Asia/Shanghai")

/** InitTime 初始化时区
 */
func InitTime() {
	var cstZone = time.FixedZone("CST", 8*3600)
	time.Local = cstZone
}

func ParseDate(dateS string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02", dateS, LOC)
}
