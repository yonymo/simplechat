package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yonymo/simplechat/pkg/code"
	gin2 "github.com/yonymo/simplechat/pkg/common/gin"
	"github.com/yonymo/simplechat/pkg/errors"
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
		gin2.HandleValidatorError(ctx, err, u.trans)
		return
	}

	if ok := store.Verify(pwdForm.CaptchaId, pwdForm.Captcha, true); !ok {
		gin2.WriteResponse(ctx, errors.WithCode(code.ErrCodeIncorrect, "验证码错误"), nil)
		return
	}

	userDTO, err := u.srv.UserSrv().MobileLogin(ctx, pwdForm.Mobile, pwdForm.PassWord)
	if err != nil {
		gin2.WriteResponse(ctx, err, nil)
		return
	}

	gin2.WriteResponse(ctx, nil, gin.H{
		"id":       userDTO.ID,
		"nickname": userDTO.Nickname,
		"token":    userDTO.Token,
		"expire":   userDTO.Expire,
	})

}
