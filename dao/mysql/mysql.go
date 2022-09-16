package mysql

import (
	"fmt"
	"web_app/settings"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

func Init(cfg *settings.MysqlConf) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
		cfg.UserName,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Dbname)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect to mysql failed ", zap.Error(err))
		return err
	}
	db.SetMaxOpenConns(cfg.MaxConn)
	db.SetMaxIdleConns(cfg.MaxIdle)
	return
}
func Close() {
	_ = db.Close()
}
