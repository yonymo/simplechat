package api

import (
	"github.com/gin-gonic/gin"

	"github.com/yonymo/simplechat/api/internal/controller/user/v1"
)

func initRouter(s *gin.Engine) {
	v1 := s.Group("/v1")
	ugroup := v1.Group("/user")
	{
		ugroup.POST("/login", user.Login)
		ugroup.POST("/register", user.Register)
	}
}
