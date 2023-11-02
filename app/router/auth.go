package router

import (
	"app/common"
	middleware "app/midellware"
	"app/modules/authentication/controller"
	"app/pkg"

	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.RouterGroup, app *pkg.AppCtx) {
	gr := r.Group("/accounts")
	gr.POST("/create", middleware.IsRole([]common.Role{common.Admin}), controller.CreateAccount(app))
	gr.POST("/password/change", middleware.IsRole([]common.Role{common.Admin}), controller.ChangePassword(app))
	gr.POST("/token/renew", middleware.IsRole([]common.Role{common.Admin}), controller.RenewToken(app))
	gr.POST("/login", controller.Login(app))
}
