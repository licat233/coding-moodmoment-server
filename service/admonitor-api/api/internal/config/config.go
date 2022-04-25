package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource   string
		MaxIdleConns int
		MaxOpenConns int
		LogMode      bool
		LogZap       string
	}

	CacheRedis cache.CacheConf

	Auth struct {
		AccessSecret string
		AccessExpire int64
		BufferTime   int64
	}
}
