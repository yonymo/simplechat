package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type PassWordLoginForm struct {
	Mobile    string `form:"mobile" json:"mobile" binding:"required,mobile"` //手机号码格式有规范可寻， 自定义validator
	PassWord  string `form:"password" json:"password" binding:"required,min=3,max=20"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
}

func (u *userServer) Login(ctx *gin.Context) {
	fmt.Println("Login is called")
	pwdForm := &PassWordLoginForm{}
	if err := ctx.ShouldBind(pwdForm); err != nil {
		ctx.JSON(http.StatusBadRequest, "参数错误")
		return
	}
	userDTO, err := u.srv.UserSrv().GetByMobile(ctx, pwdForm.Mobile)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "用户未注册")
		return
	}
	log.Println(userDTO)
}
