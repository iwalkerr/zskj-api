package db

import (
	"log"
	"sync"
	"time"
	"xframe/backend/common/cfg"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	de   *dbEngine
	once sync.Once
)

type dbEngine struct {
	master *sqlx.DB // 主数据库
	slave  *sqlx.DB // 从数据库
}

// 获取连接
func Conn(dbType ...int) *sqlx.DB {
	return instance().engine(dbType...)
}

//初始化数据操作 driver为数据库类型
func instance() *dbEngine {
	once.Do(func() {
		config := cfg.Instance()
		var db dbEngine
		// 配置主数据库
		if len(config.Database.Master) != 0 {
			db.master = db.sqlOpen(config.Database.Master)
		}
		// 配置从数据库
		if len(config.Database.Slave) != 0 {
			db.slave = db.sqlOpen(config.Database.Slave)
		}
		de = &db
	})
	return de
}

//获取操作实例 如果传入1 并且成功配置了slave 返回slave orm引擎 否则返回master orm引擎
func (db *dbEngine) engine(dbType ...int) *sqlx.DB {
	if len(dbType) > 0 && dbType[0] == 1 {
		if db.slave != nil {
			return db.slave
		}
	}
	return db.master
}

func (db *dbEngine) sqlOpen(datasource string) *sqlx.DB {
	engine, _ := sqlx.Open("mysql", datasource)
	engine.SetMaxOpenConns(1000)             // 最多打开多少个连接
	engine.SetMaxIdleConns(200)              // 设置最大的空闲连接数
	engine.SetConnMaxLifetime(time.Hour * 7) // 防止超时报错

	if err := engine.Ping(); err != nil {
		log.Fatalf("数据库连接错误: %v", err.Error())
		return nil
	}
	return engine
}
