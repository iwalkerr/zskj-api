package db

import (
	"xframe/backend/common/cfg"
	"xframe/pkg/db"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 获取连接
func Conn(dbType ...int) *sqlx.DB {
	config := cfg.Instance()
	return db.Instance(config.Database.Master, config.Database.Slave).Engine(dbType...)
}
