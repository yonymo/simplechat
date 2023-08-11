package api

import (
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	"regexp"
)

// RegisteMobile
//
//	@Description: 注册mobile的翻译器，验证器
//	@param trans
func RegisteMobile(trans ut.Translator) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", mobileValidate)
		_ = v.RegisterTranslation("mobile", trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法手机号码", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}
}

func mobileValidate(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()

	ok, _ := regexp.MatchString(`^1[3-9]\d{9}$`, mobile)
	return ok
}
