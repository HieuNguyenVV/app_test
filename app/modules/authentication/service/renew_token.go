package service

import (
	"app/common"
	"app/modules/authentication/model"
	"app/pkg/token"
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type (
	ReNewTokenRepository interface {
		GetAccountByID(ctx context.Context, id uint32) (*model.Account, error)
		GetSession(ctx context.Context, id uuid.UUID) (*model.Session, error)
		RenewSession(ctx context.Context, rfToken string, newSession *model.Session) error
	}
	RenewTokenService struct {
		repo             ReNewTokenRepository
		tokenMaker       token.Maker
		accessTokenTime  time.Duration
		refreshTokenTime time.Duration
	}
)

func NewRenewTokenService(repo ReNewTokenRepository,
	tokenMaker token.Maker,
	accessTokenTime time.Duration,
	refreshTokenTime time.Duration) *RenewTokenService {
	return &RenewTokenService{
		repo:             repo,
		tokenMaker:       tokenMaker,
		accessTokenTime:  accessTokenTime,
		refreshTokenTime: refreshTokenTime,
	}
}

func (s *RenewTokenService) RenewToken(ctx context.Context, req *model.ReNewTokenRequest) (*model.ReNewTokenResponse, error) {
	refreshPayload, err := s.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account, err := s.repo.GetAccountByID(ctx, (refreshPayload.AccountId))
	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return nil, common.ErrEntityNotFound("Account", err)
		}
		return nil, common.ErrorDB(err)
	}

	session, err := s.repo.GetSession(ctx, refreshPayload.Id)
	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return nil, common.ErrEntityNotFound("Session", err)
		}
		return nil, common.ErrorDB(err)
	}

	if account.ID != session.AccountID {
		msg := "incorrect session user"
		return nil, common.NewUnauthorized(errors.New(msg), msg, "Authorized")
	}
	if session.RefreshToken != strings.Split(req.RefreshToken, ".")[1] {
		msg := "mismatched session token"
		return nil, common.NewUnauthorized(errors.New(msg), msg, "Authorized")
	}

	if time.Now().After(session.ExpiredAt) {
		msg := "token expired"
		return nil, common.NewUnauthorized(errors.New(msg), msg, "Authorized")
	}

	accessToken, payloadAccess, err := s.tokenMaker.CreateToken(account.ID, s.accessTokenTime, account.Role)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	refreshToken, payloadRefresh, err := s.tokenMaker.CreateToken(account.ID, s.refreshTokenTime, account.Role)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	if err := s.repo.RenewSession(ctx, strings.Split(req.RefreshToken, ".")[1], &model.Session{
		ID:           payloadRefresh.Id,
		AccountID:    account.ID,
		RefreshToken: strings.Split(refreshToken, ".")[1],
		ExpiredAt:    payloadRefresh.ExpiredAt,
	}); err != nil {
		return nil, common.ErrorDB(err)
	}

	return &model.ReNewTokenResponse{
		Account: model.AccountLoginEntity{
			Role:        account.Role,
			LastLoginAt: account.LastLoginAt,
			Username:    account.Username,
		},
		AccessToken:           accessToken,
		AccessTokenExpiredAt:  payloadAccess.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiredAt: payloadRefresh.ExpiredAt,
		LastLoginAt:           account.LastLoginAt,
	}, nil
}
