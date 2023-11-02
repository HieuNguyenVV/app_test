package pkg

import (
	"app/pkg/config"
	"app/pkg/db"
	"app/pkg/env"

	"gorm.io/gorm"
)

type AppCtx struct {
	db  *gorm.DB
	evn *env.EnvType
}

func NewAppContext(db *gorm.DB, env *env.EnvType) *AppCtx {
	return &AppCtx{
		db:  db,
		evn: env,
	}
}
func (t *AppCtx) GetDBSource() *gorm.DB {
	return t.db
}

func (t *AppCtx) GetEnv() *env.EnvType {
	return t.evn
}

func App() *AppCtx {
	conf, err := config.AutoBindConfig()
	if err != nil {
		panic(err)
	}

	db, err := db.Open(conf.Postgres)
	if err != nil {
		panic(err)
	}
	env, err := env.LoadENV()
	if err != nil {
		panic(err)
	}
	app := NewAppContext(db, env)
	return app
}
