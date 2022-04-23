package initialize

import (
	"moodMoment-server/common/global"
	"moodMoment-server/common/tool"
	"moodMoment-server/config"
	"path"

	"github.com/zeromicro/go-zero/core/conf"
)

func Config() {
	global.CONFIG = new(config.Server)
	//读取yaml配置文件，映射结构体
	serverRootPath := tool.CurrentProjectRootPath()
	configFilePath := path.Join(serverRootPath, config.ProjectConfigFile)
	conf.MustLoad(configFilePath, global.CONFIG)
	global.CONFIG.ServerConfig.RootPath = serverRootPath
	global.CONFIG.Casbin.ModelPath = path.Join(serverRootPath, global.CONFIG.Casbin.ModelPath)
	global.CONFIG.Zap.Director = path.Join(serverRootPath, global.CONFIG.Zap.Director)
	//初始化静态文件存储目录
	tool.StaticServerPathHandler(&global.CONFIG.StaticConfig.Root, global.CONFIG.ServerConfig.RootPath)
}
