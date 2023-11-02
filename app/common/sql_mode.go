package common

import "time"

type SQLModel struct {
	ID       uint32    `json:"id" binding:"required" gorm:"primaryKey"`
	CreateAt time.Time `json:"createAt" binding:"required" gorm:"default:now()"`
	UpdateAt time.Time `json:"updateAt" binding:"required" gorm:"default:now()"`
	CreateBy uint32    `json:"createBy" binding:"required" gorm:"default:0"`
	UpdateBy uint32    `json:"updateBy" binding:"required" gorm:"default:0"`
}
