package model

import (
	"app/common"
	"time"
)

type Account struct {
	common.SQLModel
	Username    string      `json:"json"`
	Password    string      `json:"_"`
	Email       string      `json:"email"`
	Role        common.Role `json:"role"`
	LastLoginAt time.Time   `json:"lastLoginAt"`
}

func (Account) TableName() string {
	return "d_admins_account"
}
