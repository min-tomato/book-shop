package initialize

import (
	"fmt"

	"github.com/min-tomato/online-shop/backend/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	fmt.Println("Load configuration mysql", global.Config.Mysql.Dbname)
	InitLogger()
	global.Logger.Info("Config log oke", zap.String("oke", "success"))
	InitMysql()
	InitRedis()

	r := InitRouter()
	r.Run("8002")
}
