package options

import (
	"github.com/spf13/pflag"

	"github.com/yonymo/simplechat/pkg/errors"
)

type ServerOptions struct {
	Name string `mapstructure:"name" json:"name,omitempty"`
	Host string `mapstructure:"host" json:"host,omitempty"`
	Port int    `mapstructure:"port" json:"port,omitempty"`

	// 中间件
	Middlewares []string `mapstructure:"middlewares" json:"middlewares,omitempty"`
}

func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		Host: "127.0.0.1",
		Port: 8888,
		Name: "chat",
	}
}

func (o *ServerOptions) AddFlags(fs *pflag.FlagSet) {

	fs.StringVar(&o.Host, "server.host", o.Host, "server host address, default 127.0.0.1")
	fs.StringVar(&o.Name, "server.name", o.Name, "server name, default mxshop-user-srv")
	fs.IntVar(&o.Port, "server.port", o.Port, "server port, default 8080")

}

func (o *ServerOptions) Validate() []error {
	errs := []error{}
	if o.Host == "" || o.Port == 0 {
		errs = append(errs, errors.New("host or port is empty"))
	}
	return errs
}
