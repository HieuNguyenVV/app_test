package token

import (
	"app/common"
	"time"
)

type Maker interface {
	CreateToken(accountId uint32, duration time.Duration, role common.Role) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
