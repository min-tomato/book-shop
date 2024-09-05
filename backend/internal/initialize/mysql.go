package initialize

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/min-tomato/online-shop/backend/global"
	"go.uber.org/zap"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	m := global.Config.Mysql

	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var dbConnection = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := sql.Open("mysql", dbConnection)
	checkErrorPanic(err, "InitMysql initialization error")
	global.Logger.Info("Initializing MySQL successfully")

	global.Mdb = db
	SetPool(db)
}

func SetPool(db *sql.DB) {
	m := global.Config.Mysql
	db.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	db.SetMaxOpenConns(m.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(m.ConnMaxLifeTime))
}
