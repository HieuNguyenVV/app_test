package model

import "time"

type ReNewTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

type ReNewTokenResponse struct {
	Account               AccountLoginEntity `json:"account" binding:"required"`
	AccessToken           string             `json:"accessToken" binding:"required"`
	AccessTokenExpiredAt  time.Time          `json:"accessTokenExpiredAt" binding:"required"`
	RefreshToken          string             `json:"refreshToken" binding:"required"`
	RefreshTokenExpiredAt time.Time          `json:"refreshTokenExpiredAt" binding:"required"`
	LastLoginAt           time.Time          `json:"lastLoginAt" binding:"required"`
}
