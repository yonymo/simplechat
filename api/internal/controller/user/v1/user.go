package user

import "github.com/yonymo/simplechat/api/internal/service"

type userServer struct {
	srv service.ServiceFactory
}

func NewUserControl(srv service.ServiceFactory) *userServer {
	return &userServer{srv: srv}
}
