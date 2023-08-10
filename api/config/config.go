package config

import "github.com/yonymo/simplechat/pkg/options"

type Config struct {
	Mysql *options.MySQLOptions `json:"mysql" mapstructure:"mysql"`
}

func NewConfig() *Config {
	return &Config{Mysql: options.NewMySQLOptions()}
}
