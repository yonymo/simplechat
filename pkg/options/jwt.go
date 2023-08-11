package options

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/spf13/pflag"
	"time"
)

type JwtOptions struct {
	Realm      string        `json:"realm" mapstructure:"realm"`
	Key        string        `json:"key" mapstructure:"key"`
	Timeout    time.Duration `json:"timeout" mapstructure:"timeout"`
	MaxRefresh time.Duration `json:"max-refresh" mapstructure:"max-refresh"`
}

func NewJwtOptions() *JwtOptions {
	return &JwtOptions{
		Realm:      "chat",
		Key:        "chat-key123453asgfv",
		Timeout:    24 * time.Hour,
		MaxRefresh: 24 * time.Hour,
	}
}

func (s *JwtOptions) Validate() []error {
	var errs []error
	if !govalidator.StringLength(s.Key, "6", "32") {
		errs = append(errs, fmt.Errorf("secret key length must between 6 and 32."))
	}
	return errs
}

func (s *JwtOptions) AddFlags(fs *pflag.FlagSet) {
	if fs == nil {
		return
	}

	fs.StringVar(&s.Realm, "jwt.realm", s.Realm, "jwt realm")
	fs.StringVar(&s.Key, "jwt.key", s.Key, "jwt secret key.")
	fs.DurationVar(&s.Timeout, "jwt.timeout", s.Timeout, "jwt token timeout")
	fs.DurationVar(&s.MaxRefresh, "jwt.max-refresh", s.MaxRefresh, "refresh token to max refresh time.")
}
