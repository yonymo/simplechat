package api

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/yonymo/simplechat/api/config"
	fdb "github.com/yonymo/simplechat/api/internal/data/friend/v1/db"
	udb "github.com/yonymo/simplechat/api/internal/data/user/v1/db"
	"github.com/yonymo/simplechat/api/internal/service"
	"github.com/yonymo/simplechat/pkg/common/db"

	"github.com/yonymo/simplechat/api/internal/controller/friend/v1"
	"github.com/yonymo/simplechat/api/internal/controller/user/v1"
)

func initRouter(s *gin.Engine, cfg *config.Config, trans ut.Translator) {
	dbins, err := db.GetDBInstance(cfg.Mysql)
	if err != nil {
		panic(err)
	}
	v1 := s.Group("/v1")

	userData := udb.NewUserData(dbins)
	friendData := fdb.NewFriendData(dbins)
	srvParam := &service.FactoryParam{
		UserData:   userData,
		JwtOps:     cfg.Jwt,
		FriendData: friendData,
	}
	srvFact := service.NewSrvFactory(srvParam)

	ugroup := v1.Group("/user")
	{
		userServer := user.NewUserControl(srvFact, trans)
		ugroup.POST("/login", userServer.Login)
		ugroup.POST("/register", userServer.Register)
	}

	baseGroup := v1.Group("/base")
	{
		baseGroup.GET("/captcha", user.GetCaptcha)
	}

	jwtAuth := newJWTAuth(cfg.Jwt)

	friendGroup := v1.Group("/friend")
	{
		friendGroup.Use(jwtAuth.AuthFunc())
		friendServer := friend.NewFriendControl(srvFact)
		friendGroup.POST("/add_friend", friendServer.AddFriend)
		friendGroup.GET("/list", friendServer.List)
	}
}
