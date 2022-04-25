package main

import (
	"flag"
	"fmt"
	"moodMoment-server/common/errorx"
	"moodMoment-server/common/initialize"
	"moodMoment-server/service/admonitor-api/api/internal/config"
	"moodMoment-server/service/admonitor-api/api/internal/handler"
	"moodMoment-server/service/admonitor-api/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/admonitor-api.yaml", "the config file")

func initLogger() {
	c := logx.LogConf{
		ServiceName:         "admonitor-api",
		Mode:                "file",
		Encoding:            "plain",
		TimeFormat:          "",
		Path:                "logxs",
		Level:               "error",
		Compress:            true,
		KeepDays:            3,
		StackCooldownMillis: 200,
	}
	logx.MustSetup(c)
}

func main() {
	flag.Parse()

	initialize.Initialize()
	// 初始化logx日志，必须放在前面这个地方
	initLogger()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandler(errorx.ErrorToErrorX)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
