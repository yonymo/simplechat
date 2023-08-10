package service

import (
	du "github.com/yonymo/simplechat/api/internal/data/user/v1"
	"github.com/yonymo/simplechat/api/internal/service/user/v1"
)

type ServiceFactory interface {
	UserSrv() user.IUserSrv
}

type serviceFac struct {
	userData du.IUserData
}

func (s *serviceFac) UserSrv() user.IUserSrv {
	return user.NewUserSrv(s.userData)
}

var _ ServiceFactory = &serviceFac{}

func NewSrvFactory(ud du.IUserData) ServiceFactory {
	return &serviceFac{userData: ud}
}
