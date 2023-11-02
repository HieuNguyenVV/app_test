package controller

import (
	"app/common"
	"app/modules/authentication/model"
	"app/modules/authentication/repository"
	"app/modules/authentication/service"
	"app/pkg"
	"app/pkg/token"

	"github.com/gin-gonic/gin"
)

func ChangePassword(app *pkg.AppCtx) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := model.ChangePasswordRequest{}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		repo := repository.NewSQLRepository(app.GetDBSource())
		s := service.NewChangePasswordService(repo)

		account := ctx.MustGet(token.CurrentUser).(*token.Payload)
		if err := s.ChangePassword(ctx.Request.Context(), &req, account.AccountId); err != nil {
			panic(err)
		}
	}
}
