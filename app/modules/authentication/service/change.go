package service

import (
	"app/common"
	"app/helper"
	"app/modules/authentication/model"
	"context"
	"errors"
)

type (
	ChangePasswordRepository interface {
		ChangePassword(ctx context.Context, hashedPass, email string, accountId uint32) error
		GetAccountByID(ctx context.Context, id uint32) (*model.Account, error)
	}
	ChangePasswordService struct {
		repo ChangePasswordRepository
	}
)

func NewChangePasswordService(repo ChangePasswordRepository) *ChangePasswordService {
	return &ChangePasswordService{
		repo: repo,
	}
}

func (s *ChangePasswordService) ChangePassword(ctx context.Context, req *model.ChangePasswordRequest, accountID uint32) error {
	if req.ConfirmPassword != req.NewPassword {
		err := errors.New("Confirm password fail")
		return common.ErrorWrongConfirmPassword(err)
	}
	account, err := s.repo.GetAccountByID(ctx, accountID)
	if err != nil {
		return common.ErrorDB(err)
	}

	if err := helper.CheckPassword(req.OldPassword, account.Password); err != nil {
		err := errors.New("Check password fail")
		return common.ErrWrongCurrentPassword(err)
	}
	hashedPass, err := helper.HashPassword(req.NewPassword)
	if err != nil {
		common.ErrInternal(err)
	}
	if err := s.repo.ChangePassword(ctx, hashedPass, account.Email, accountID); err != nil {
		return common.ErrorDB(err)
	}
	return nil
}
