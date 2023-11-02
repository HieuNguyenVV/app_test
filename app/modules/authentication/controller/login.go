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

func Login(app *pkg.AppCtx) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := model.AccountLoginRequest{}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
			//panic(common.ErrInvalidRequest(err))
		}
		tokenMaker, err := token.NewJWTMaker(app.GetEnv().SecretKey)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.ErrInternal(err))
			return
			//panic(common.ErrInternal(err))
		}
		repo := repository.NewSQLRepository(app.GetDBSource())
		s := service.NewLoginService(repo, tokenMaker, app.GetEnv().AccessTokenTime, app.GetEnv().RefreshTokenTime)
		resp, err := s.Login(ctx.Request.Context(), &req)
		if err != nil {
			if err.(*common.AppError).StatusCode == 400 {
				ctx.JSON(http.StatusBadRequest, common.ErrorDB(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, common.ErrInternal(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
