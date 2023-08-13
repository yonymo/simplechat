package auth

import (
	"github.com/yonymo/simplechat/pkg/code"
	gin2 "github.com/yonymo/simplechat/pkg/common/gin"
	"github.com/yonymo/simplechat/pkg/errors"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/yonymo/simplechat/pkg/middleware"
)

const authHeaderCount = 2

// AutoStrategy defines authentication strategy which can automatically choose between Basic and Bearer
// according `Authorization` header.
type AutoStrategy struct {
	jwt JWTStrategy
}

var _ middleware.AuthStrategy = &AutoStrategy{}

// NewAutoStrategy create auto strategy with basic strategy and jwt strategy.
func NewAutoStrategy(jwt JWTStrategy) AutoStrategy {
	return AutoStrategy{
		jwt: jwt,
	}
}

// AuthFunc defines auto strategy as the gin authentication middleware.
func (a AutoStrategy) AuthFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		operator := middleware.AuthOperator{}
		authHeader := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(authHeader) != authHeaderCount {
			gin2.WriteResponse(
				c,
				errors.WithCode(code.ErrInvalidAuthHeader, "Authorization header format is wrong."),
				nil,
			)
			c.Abort()

			return
		}

		switch authHeader[0] {
		case "Bearer":
			operator.SetStrategy(a.jwt)
			// a.JWT.MiddlewareFunc()(c)
		default:
			gin2.WriteResponse(c, errors.WithCode(code.ErrSignatureInvalid, "unrecognized Authorization header."), nil)
			c.Abort()

			return
		}

		operator.AuthFunc()(c)

		c.Next()
	}
}
