package initialize

import (
	"github.com/min-tomato/online-shop/backend/global"
	"github.com/min-tomato/online-shop/backend/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
