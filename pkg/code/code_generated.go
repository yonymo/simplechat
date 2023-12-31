// Code generated by "codegen -type=int"; DO NOT EDIT.

package code

// init register error codes defines in this source code to `github.com/marmotedu/errors`
func init() {
	register(ErrParam, 400, "Param error")
	register(ErrTokenCreate, 400, "Token create failed")
	register(ErrUserNotFound, 404, "User not found")
	register(ErrUserAlreadyExists, 400, "User already exists")
	register(ErrUserPasswordIncorrect, 400, "User password incorrect")
	register(ErrSmsSend, 400, "Send sms error")
	register(ErrCodeNotExist, 400, "Sms code incorrect or expired")
	register(ErrCodeIncorrect, 400, "Verify code incorrect")
}
