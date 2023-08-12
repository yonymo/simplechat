package service

import (
	df "github.com/yonymo/simplechat/api/internal/data/friend/v1"
	du "github.com/yonymo/simplechat/api/internal/data/user/v1"
	"github.com/yonymo/simplechat/api/internal/service/friend/v1"
	"github.com/yonymo/simplechat/api/internal/service/user/v1"
	"github.com/yonymo/simplechat/pkg/options"
	"sync"
)

type ServiceFactory interface {
	UserSrv() user.IUserSrv
	FriendSrv() friend.IFriendSrv
}

var (
	friendSrv friend.IFriendSrv
	userSrv   user.IUserSrv
	srvFact   ServiceFactory
	once      sync.Once
)

type serviceFac struct {
}

func (s *serviceFac) FriendSrv() friend.IFriendSrv {
	return friendSrv
}

func (s *serviceFac) UserSrv() user.IUserSrv {
	return userSrv
}

var _ ServiceFactory = &serviceFac{}

type FactoryParam struct {
	UserData   du.IUserData
	JwtOps     *options.JwtOptions
	FriendData df.IFriendData
}

func NewSrvFactory(param *FactoryParam) ServiceFactory {
	once.Do(func() {
		friendSrv = friend.NewFriendSrv(param.FriendData)
		userSrv = user.NewUserSrv(param.UserData, param.JwtOps)
		srvFact = &serviceFac{}
	})

	return srvFact
}
