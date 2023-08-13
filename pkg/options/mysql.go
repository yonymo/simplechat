package options

import (
	"github.com/spf13/pflag"
	"github.com/yonymo/simplechat/pkg/errors"
	"time"
)

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
		Database:              "simple-chat",
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		MaxConnectionLifetime: time.Duration(10) * time.Second,
		LogLevel:              4, // silent
	}
}

func (o *MySQLOptions) Validate() []error {
	errs := []error{}
	if o.Database == "" {
		errs = append(errs, errors.New("Database is empty"))
	}
	return errs
}

func (o *MySQLOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Host, "mysql.host", o.Host, "Mysql Service host address.")
	fs.StringVar(&o.Port, "mysql.port", o.Port, "Mysql Service port.")
	fs.StringVar(&o.UserName, "mysql.username", o.UserName, "Mysql Service username.")
	fs.StringVar(&o.Password, "mysql.password", o.Password, "Mysql Service password.")
	fs.StringVar(&o.Database, "mysql.database", o.Database, "Mysql Service database.")
	fs.IntVar(&o.MaxIdleConnections, "mysql.max-idle-connections", o.MaxIdleConnections, "Mysql Service max idle connections.")
	fs.IntVar(&o.MaxOpenConnections, "mysql.max-open-connections", o.MaxOpenConnections, "Mysql Service max open connections.")
	fs.DurationVar(&o.MaxConnectionLifetime, "mysql.max-connection-life-time", o.MaxConnectionLifetime, "Mysql Service max connection life time.")
	fs.IntVar(&o.LogLevel, "mysql.log-level", o.LogLevel, "Mysql Service log level.")

}
