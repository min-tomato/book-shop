package global

import (
	"github.com/min-tomato/online-shop/backend/pkg/logger"
	"github.com/min-tomato/online-shop/backend/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
