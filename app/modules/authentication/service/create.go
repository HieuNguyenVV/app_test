package service

import (
	"app/common"
	"app/helper"
	"app/modules/authentication/model"
	"context"
	"errors"
	"time"
)

type (
	CreateAccountRepository interface {
		CreateAccount(ctx context.Context, data *model.Account) error
		CheckExistByEmail(ctx context.Context, email string) (bool, error)
	}
	CreateAccountService struct {
		repo CreateAccountRepository
	}
)

func NewCreateAccountService(repo CreateAccountRepository) *CreateAccountService {
	return &CreateAccountService{
		repo: repo,
	}
}

func (s *CreateAccountService) CreateAccount(ctx context.Context, req *model.CreateAccountRequest) error {
	exist, err := s.repo.CheckExistByEmail(ctx, req.Email)
	if err != nil {
		return common.ErrorDB(err)
	}
	if exist {
		msg := "email existed"
		err := errors.New(msg)
		return common.NewErrorResponse(err, msg, err.Error(), "EMAIL_EXISTED")
	}
	hashedPass, err := helper.HashPassword(req.Password)
	if err != nil {
		return common.ErrInternal(err)
	}
	if err := s.repo.CreateAccount(ctx, &model.Account{
		Username:    req.Username,
		Password:    hashedPass,
		Email:       req.Email,
		Role:        common.Admin,
		LastLoginAt: time.Now(),
	}); err != nil {
		return common.ErrorDB(err)
	}
	return nil
}
