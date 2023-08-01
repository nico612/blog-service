package global

import (
	"github.com/nico612/blog-service/pkg/logger"
	"github.com/nico612/blog-service/pkg/setting"
)

// 全局配置文件信息

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
