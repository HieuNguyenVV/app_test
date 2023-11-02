package model

type CreateAccountRequest struct {
	Username string `json:"username" binding:"required,lte=255,gte=6"`
	Password string `json:"password" binding:"required,lte=255,gte=6"`
	Email    string `json:"email" binding:"required,email"`
}

type AccountLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,lte=255,gte=6"`
}
