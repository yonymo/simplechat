package user

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/yonymo/simplechat/api/internal/service"
)

type userServer struct {
	srv   service.ServiceFactory
	trans ut.Translator
}

func NewUserControl(srv service.ServiceFactory, trans ut.Translator) *userServer {
	return &userServer{srv: srv, trans: trans}
}
