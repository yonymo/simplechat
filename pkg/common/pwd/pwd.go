package pwd

import (
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"strings"
)

var options = &password.Options{16, 100, 32, sha512.New}

func EncodePwd(passwd string) string {
	if len(passwd) == 0 {
		return ""
	}
	//密码加密
	salt, encodedPwd := password.Encode(passwd, options)
	return fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
}

func VerifyPwd(passwd, encryptPwd string) bool {
	//校验密码
	passwordInfo := strings.Split(encryptPwd, "$")
	return password.Verify(passwd, passwordInfo[2], passwordInfo[3], options)
}
