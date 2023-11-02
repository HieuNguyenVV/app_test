package controller

import (
	"app/common"
	"app/modules/authentication/model"
	"app/modules/authentication/repository"
	"app/modules/authentication/service"
	"app/pkg"
	"app/pkg/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenewToken(app *pkg.AppCtx) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := model.ReNewTokenRequest{}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		tokenMaker, err := token.NewJWTMaker(app.GetEnv().SecretKey)
		if err != nil {
			panic(common.ErrInternal(err))
		}
		repo := repository.NewSQLRepository(app.GetDBSource())
		s := service.NewRenewTokenService(repo, tokenMaker, app.GetEnv().AccessTokenTime, app.GetEnv().RefreshTokenTime)

		resp, err := s.RenewToken(ctx.Request.Context(), &req)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
