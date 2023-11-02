package token

import (
	"app/common"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrorInvalidToken = errors.New("invalid token")
	ErrorExpiredToken = errors.New("error expired token")
)

const CurrentUser = "authorization_payload"

type Payload struct {
	Id        uuid.UUID   `json:"id"`
	AccountId uint32      `json:"accountId"`
	Role      common.Role `json:"role"`
	ExpiredAt time.Time   `json:"expiredAt"`
	IssueAt   time.Time   `json:"issueAt"`
}

func NewPayload(accountId uint32, duration time.Duration, role common.Role) (*Payload, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return &Payload{
		Id:        id,
		AccountId: accountId,
		Role:      role,
		ExpiredAt: time.Now().Add(duration),
		IssueAt:   time.Now(),
	}, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrorExpiredToken
	}
	return nil
}

func (payload *Payload) GetAccountId() uint32 {
	return payload.AccountId
}
