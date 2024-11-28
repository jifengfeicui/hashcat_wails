package global

import (
	"context"
	"os"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"wails_hashcat/config"
)

//var DB *gorm.DB

var (
	Dir         string
	SugarLogger *zap.SugaredLogger
	StartupCtx  context.Context
	CONFIG      config.System // 系统配置信息
	DB          *gorm.DB
)

func init() {
	Dir, _ = os.Getwd()

}
