package main

import (
	middleware "app/midellware"
	"app/pkg"
	"app/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// conf, err := config.AutoBindConfig()
	// if err != nil {
	// 	panic(err)
	// }

	// db, err := db.Open(conf.Postgres)
	// if err != nil {
	// 	panic(err)
	// }
	// env, err := env.LoadENV()
	// if err != nil {
	// 	panic(err)
	// }
	app := pkg.App()
	r := gin.Default()
	gr := r.Group("/api/v1")
	gr.Use(middleware.Auth(app.GetEnv().SecretKey))
	//app := pkg.NewAppContext(db, env)

	router.AuthRouter(gr, app)
	r.Run(":8080")
}
