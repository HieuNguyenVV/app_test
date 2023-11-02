package service

import (
	"app/common"
	"app/helper"
	"app/modules/authentication/model"
	"app/pkg/token"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
)

type (
	LoginRepository interface {
		CreateSession(ctx context.Context, data *model.Session) error
		CheckExistByEmail(ctx context.Context, email string) (bool, error)
		GetAccountByEmail(ctx context.Context, email string) (*model.Account, error)
	}
	LoginService struct {
		repo             LoginRepository
		tokenMaker       token.Maker
		accessTokenTime  time.Duration
		refreshTokenTime time.Duration
	}
)

func NewLoginService(repo LoginRepository,
	tokenMaker token.Maker,
	accessTokenTime time.Duration,
	refreshTokenTime time.Duration) *LoginService {
	return &LoginService{
		repo:             repo,
		tokenMaker:       tokenMaker,
		accessTokenTime:  accessTokenTime,
		refreshTokenTime: refreshTokenTime,
	}
}

func (s *LoginService) Login(ctx context.Context, req *model.AccountLoginRequest) (*model.AccountLoginResponse, error) {
	exist, err := s.repo.CheckExistByEmail(ctx, req.Email)
	if err != nil {
		return nil, common.ErrorDB(err)
	}
	if !exist {
		msg := "email does not exist"
		err := errors.New(msg)
		return nil, common.NewErrorResponse(err, msg, err.Error(), "EMAIL_NOT_EXIST")
	}
	fmt.Println(req.Password)
	account, err := s.repo.GetAccountByEmail(ctx, req.Email)

	if err != nil {
		return nil, common.ErrorDB(err)
	}

	if err := helper.CheckPassword(req.Password, account.Password); err != nil {
		return nil, common.ErrInternal(err)
	}

	accessToken, accessPayload, err := s.tokenMaker.CreateToken(account.ID, s.accessTokenTime, account.Role)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	if err := helper.CheckPassword(req.Password, account.Password); err != nil {
		return nil, common.ErrInternal(err)
	}

	refreshToken, refreshPayload, err := s.tokenMaker.CreateToken(account.ID, s.refreshTokenTime, account.Role)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	if err := helper.CheckPassword(req.Password, account.Password); err != nil {
		return nil, common.ErrInternal(err)
	}

	if err := s.repo.CreateSession(ctx, &model.Session{
		ID:           refreshPayload.Id,
		AccountID:    account.ID,
		RefreshToken: strings.Split(refreshToken, ".")[1],
		ExpiredAt:    accessPayload.ExpiredAt,
	}); err != nil {
		return nil, common.ErrInternal(err)
	}
	return &model.AccountLoginResponse{
		Account: model.AccountLoginEntity{
			Role:        account.Role,
			LastLoginAt: account.LastLoginAt,
			Username:    account.Username,
		},
		AccessToken:           accessToken,
		AccessTokenExpiredAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiredAt: refreshPayload.ExpiredAt,
		LastLoginAt:           account.LastLoginAt,
	}, nil
}
