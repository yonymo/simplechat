package friend

import (
	"github.com/yonymo/simplechat/api/internal/service"
)

type friendServer struct {
	srv service.ServiceFactory
}

func NewFriendControl(srv service.ServiceFactory) *friendServer {
	return &friendServer{srv: srv}
}
