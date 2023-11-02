package db

import (
	"app/modules/authentication/model"
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (t Postgres) getSourceDB() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
		t.Host, t.Username, t.Password, t.DBName, t.Port)
}

func Open(conf Postgres) (*gorm.DB, error) {
	strCon := conf.getSourceDB()
	db, err := gorm.Open(postgres.Open(strCon))
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(model.Account{}, model.Session{})
	return db.Debug(), nil
}
