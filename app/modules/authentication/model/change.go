package model

type ChangePasswordRequest struct {
	OldPassword     string `json:"oldPassword" binding:"required,min=6"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}
