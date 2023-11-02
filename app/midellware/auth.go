package middleware

import (
	"app/common"
	"app/pkg/token"
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type PublicURLs []string

var publicURLs = PublicURLs{
	"/api/v1/accounts/login",
}

const (
	AuthorizationHeaderKey    = "Authorization"
	AuthorizationHeaderBearer = "Bearer"
)

func (url *PublicURLs) Contain(str string) bool {
	for _, v := range *url {
		if v == str {
			fmt.Println("OK")
			return true
		}
	}
	return false
}

func Auth(secretKey string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenMaker, err := token.NewJWTMaker(secretKey)
		if err != nil {
			panic(err)
		}
		fmt.Println(ctx.FullPath())
		if publicURLs.Contain(ctx.FullPath()) {
			ctx.Next()
			return
		}

		tokenJwt := ctx.GetHeader(AuthorizationHeaderKey)
		fmt.Println(tokenJwt)
		if len(tokenJwt) == 0 {
			message := "authorization header is not provided"
			panic(common.NewUnauthorized(errors.New(message), message, "Authorized"))
		}

		fields := strings.Fields(tokenJwt)

		if len(fields) < 2 {
			message := "invalid authorization header format"
			panic(common.NewUnauthorized(errors.New(message), message, "Unauthorized"))
		}

		authorizationType := fields[0]
		if authorizationType != AuthorizationHeaderBearer {
			message := fmt.Sprintf("unsupported authorization type %s", authorizationType)
			panic(common.NewUnauthorized(errors.New(message), message, "Unauthorized"))
		}

		accessToken := fields[1]

		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			panic(common.NewUnauthorized(err, err.Error(), "Unauthorized"))
		}

		ctx.Set(token.CurrentUser, payload)
		ctx.Next()
	}
}
