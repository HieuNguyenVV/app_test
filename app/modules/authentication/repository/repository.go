package repository

import (
	"gorm.io/gorm"
)

type SQLRepository struct {
	*gorm.DB
}

func NewSQLRepository(db *gorm.DB) *SQLRepository {
	return &SQLRepository{
		DB: db,
	}
}
