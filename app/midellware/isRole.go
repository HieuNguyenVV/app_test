package middleware

import (
	"app/common"
	"app/pkg/token"
	"errors"

	"github.com/gin-gonic/gin"
)

func IsRole(role []common.Role) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		account := ctx.MustGet(token.CurrentUser).(*token.Payload)
		if !Contain(role, account.Role) {
			message := "Forbidden"
			panic(common.NewForbidden(errors.New(message), message, message))
		}
		ctx.Next()
	}
}
func Contain(s []common.Role, role common.Role) bool {
	for _, v := range s {
		if v == role {
			return true
		}
	}
	return false
}
