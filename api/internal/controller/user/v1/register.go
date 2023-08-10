package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (u *userServer) Register(ctx *gin.Context) {
	fmt.Println("Register is called")
}
