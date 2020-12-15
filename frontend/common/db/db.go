package db

import (
	"xframe/frontend/config"
	"xframe/pkg/db"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 获取连接
func Conn(dbType ...int) *sqlx.DB {
	return db.Instance(config.MysqlMaster, config.MysqlSlave).Engine(dbType...)
}
