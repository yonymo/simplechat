package service

import (
	du "github.com/yonymo/simplechat/api/internal/data/user/v1"
	"github.com/yonymo/simplechat/api/internal/service/user/v1"
	"github.com/yonymo/simplechat/pkg/options"
)

type ServiceFactory interface {
	UserSrv() user.IUserSrv
}

type serviceFac struct {
	userData du.IUserData
	jwtOps   *options.JwtOptions
}

func (s *serviceFac) UserSrv() user.IUserSrv {
	return user.NewUserSrv(s.userData, s.jwtOps)
}

var _ ServiceFactory = &serviceFac{}

func NewSrvFactory(ud du.IUserData, jwtOps *options.JwtOptions) ServiceFactory {
	return &serviceFac{userData: ud, jwtOps: jwtOps}
}
