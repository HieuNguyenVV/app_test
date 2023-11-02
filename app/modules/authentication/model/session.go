package model

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID  `json:"id"`
	AccountID    uint32     `json:"accountId"`
	RefreshToken string     `json:"refreshToken"`
	ExpiredAt    time.Time  `json:"expiredAt"`
	CreatedBy    uint32     `json:"createdBy" gorm:"not null; default:0"`
	UpdatedBy    uint32     `json:"updatedBy" gorm:"not null; default:0"`
	CreatedAt    *time.Time `json:"createdAt" gorm:"not null; default:now()"`
	UpdatedAt    *time.Time `json:"updatedAt" gorm:"not null; default:now()"`
}

func (Session) TableName() string {
	return "d_admins_token"
}
