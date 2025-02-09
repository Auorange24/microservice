package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	MySqlConfig MySqlConfig
}

// MySQL的配置
type MySqlConfig struct {
	DataSource string
	ConnectTimeOut int64
}