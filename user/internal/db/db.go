package db

import (
	"user/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"context"
	"time"
)

func NewMySql(MySqlConfig config.MySqlConfig) sqlx.SqlConn{
	// 创建mysql connection
	mysql := sqlx.NewMysql(MySqlConfig.DataSource)
	// RawDB 是一个操作对象
	/*
	connection 和 RawDB 的区别
	*/
	db, err := mysql.RawDB()
	if err != nil {
		panic(err)
	}
	/*
	创建一个超时context 测试能否连接到mysql
	*/
	cxt, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(MySqlConfig.ConnectTimeOut))
	defer cancel()
	err = db.PingContext(cxt)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	return mysql
}