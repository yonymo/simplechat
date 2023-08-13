package api

import (
	ginjwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/yonymo/simplechat/pkg/middleware"
	"github.com/yonymo/simplechat/pkg/middleware/auth"
	"github.com/yonymo/simplechat/pkg/options"
)

func newJWTAuth(opts *options.JwtOptions) middleware.AuthStrategy {
	gjwt, _ := ginjwt.New(&ginjwt.GinJWTMiddleware{
		Realm:            opts.Realm,
		SigningAlgorithm: "HS256",
		Key:              []byte(opts.Key),
		Timeout:          opts.Timeout,
		MaxRefresh:       opts.MaxRefresh,
		LogoutResponse: func(c *gin.Context, code int) {
			c.JSON(code, nil)
		},
		IdentityHandler: claimHandlerFun,
		IdentityKey:     middleware.KeyUserID,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
	})

	return auth.NewJWTStrategy(*gjwt)
}

func claimHandlerFun(c *gin.Context) interface{} {
	claims := ginjwt.ExtractClaims(c)
	c.Set(middleware.KeyUserID, claims[middleware.KeyUserID])
	return claims[ginjwt.IdentityKey]
}
