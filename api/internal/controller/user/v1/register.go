package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	gin2 "github.com/yonymo/simplechat/pkg/common/gin"
	"net/http"
)

type RegisterForm struct {
	Mobile    string `form:"mobile" json:"mobile" binding:"required,mobile"` //手机号码格式有规范可寻， 自定义validator
	PassWord  string `form:"password" json:"password" binding:"required,min=3,max=20"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
}

func (u *userServer) Register(ctx *gin.Context) {
	fmt.Println("Register is called")

	regForm := &RegisterForm{}

	if err := ctx.ShouldBind(regForm); err != nil {
		gin2.HandleValidatorError(ctx, err, u.trans)
		return
	}

	// 验证码验证
	if !store.Verify(regForm.CaptchaId, regForm.Captcha, true) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"captcha": "验证码错误",
		})
		return
	}

	dto, err := u.srv.UserSrv().Register(ctx, regForm.Mobile, regForm.PassWord)
	if err != nil {
		gin2.WriteResponse(ctx, err, nil)
		return
	}

	gin2.WriteResponse(ctx, nil, gin.H{
		"id":       dto.ID,
		"nickname": dto.Nickname,
		"token":    dto.Token,
		"expire":   dto.Expire,
	})
}
