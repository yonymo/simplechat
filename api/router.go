package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yonymo/simplechat/api/config"
	"github.com/yonymo/simplechat/api/internal/data/user/v1/db"
	"github.com/yonymo/simplechat/api/internal/service"

	"github.com/yonymo/simplechat/api/internal/controller/user/v1"
)

func initRouter(s *gin.Engine, cfg *config.Config) {
	dbins, err := db.GetDBInstance(cfg.Mysql)
	if err != nil {
		panic(err)
	}
	v1 := s.Group("/v1")
	ugroup := v1.Group("/user")
	userData := db.NewUserData(dbins)
	srvFact := service.NewSrvFactory(userData)
	userServer := user.NewUserControl(srvFact)
	{
		ugroup.POST("/login", userServer.Login)
		ugroup.POST("/register", userServer.Register)
	}
}
