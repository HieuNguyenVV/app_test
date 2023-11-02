package controller

import (
	"app/common"
	"app/modules/authentication/model"
	"app/modules/authentication/repository"
	"app/modules/authentication/service"
	"app/pkg"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAccount(app *pkg.AppCtx) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := model.CreateAccountRequest{}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			//panic(common.ErrInvalidRequest(err))
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		repo := repository.NewSQLRepository(app.GetDBSource())

		s := service.NewCreateAccountService(repo)
		err := s.CreateAccount(ctx.Request.Context(), &req)
		if err != nil {
			//panic(err)
			fmt.Println(err)
			if err.(*common.AppError).StatusCode == 400 {
				ctx.JSON(http.StatusBadRequest, common.ErrorDB(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, common.ErrInternal(err))
		}
	}

}
