package model

import (
	"app/common"
	"time"
)

type AccountLoginEntity struct {
	Role        common.Role `json:"role" binding:"required,role" enums:"1,99"`
	LastLoginAt time.Time   `json:"lastLoginAt" binding:"required"`
	Username    string      `json:"username" binding:"required"`
} // @name AccountLoginEntity

type AccountLoginResponse struct {
	Account               AccountLoginEntity `json:"account" binding:"required"`
	AccessToken           string             `json:"accessToken" binding:"required"`
	AccessTokenExpiredAt  time.Time          `json:"accessTokenExpiredAt" binding:"required"`
	RefreshToken          string             `json:"refreshToken" binding:"required"`
	RefreshTokenExpiredAt time.Time          `json:"refreshTokenExpiredAt" binding:"required"`
	LastLoginAt           time.Time          `json:"lastLoginAt" binding:"required"`
} // @name AccountLoginResponse
