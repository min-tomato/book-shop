package global

import (
	"database/sql"

	"github.com/min-tomato/online-shop/backend/pkg/logger"
	"github.com/min-tomato/online-shop/backend/pkg/setting"
	"github.com/redis/go-redis/v9"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *sql.DB
	Rdb    *redis.Client
)
