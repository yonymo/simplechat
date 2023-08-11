package config

import (
	cliflag "github.com/yonymo/simplechat/pkg/common/cli/flag"
	"github.com/yonymo/simplechat/pkg/options"
)

type Config struct {
	Mysql  *options.MySQLOptions  `json:"mysql" mapstructure:"mysql"`
	Server *options.ServerOptions `json:"server" mapstructure:"server"`
	Jwt    *options.JwtOptions    `json:"jwt" mapstructure:"jwt"`
}

func NewConfig() *Config {
	return &Config{
		Mysql:  options.NewMySQLOptions(),
		Server: options.NewServerOptions(),
		Jwt:    options.NewJwtOptions(),
	}
}

func (c *Config) Flags() (fss cliflag.NamedFlagSets) {
	c.Mysql.AddFlags(fss.FlagSet("mysql"))
	c.Server.AddFlags(fss.FlagSet("server"))
	c.Jwt.AddFlags(fss.FlagSet("jwt"))

	return
}

func (c *Config) Validate() []error {
	var errs []error
	errs = append(errs, c.Mysql.Validate()...)
	errs = append(errs, c.Server.Validate()...)
	errs = append(errs, c.Jwt.Validate()...)

	return errs
}
