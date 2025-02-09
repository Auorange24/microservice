package svc

import (
	"user/internal/config"
	"user/internal/db"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Conn sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := db.NewMySql(c.MySqlConfig)
	return &ServiceContext{
		Config: c,
		Conn: sqlConn,
	}
}
