package options

import "time"

type MySQLOptions struct {
	Host                  string        `json:"host,omitempty" mapstructure:"host"`
	Port                  string        `json:"port,omitempty" mapstructure:"port"`
	UserName              string        `json:"username,omitempty" mapstructure:"username"`
	Password              string        `json:"password,omitempty" mapstructure:"password"`
	Database              string        `json:"database,omitempty" mapstructure:"database"`
	MaxIdleConnections    int           `json:"max-idle-connections,omitempty" mapstructure:"max-idle-connections"`
	MaxOpenConnections    int           `json:"max-open-connections,omitempty" mapstructure:"max-open-connections"`
	MaxConnectionLifetime time.Duration `json:"max-connection-life-time,omitempty" mapstructure:"max-connection-life-time"`
	LogLevel              int           `json:"log-level,omitempty" mapstructure:"log-level"`
}

func NewMySQLOptions() *MySQLOptions {
	return &MySQLOptions{
		Host:                  "127.0.0.1",
		Port:                  "3306",
		UserName:              "root",
		Password:              "123456",
		Database:              "",
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		MaxConnectionLifetime: time.Duration(10) * time.Second,
		LogLevel:              1, // silent
	}
}
