package token

import (
	"app/common"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTMaker struct {
	secretKey string
}

const (
	maxSizeSecretKey = 32
)

func NewJWTMaker(secretKey string) (*JWTMaker, error) {
	if len(secretKey) < maxSizeSecretKey {
		return nil, fmt.Errorf("the size of secret key key must be must be great than %s", maxSizeSecretKey)
	}
	return &JWTMaker{
		secretKey: secretKey,
	}, nil
}

func (t *JWTMaker) CreateToken(accountId uint32, duration time.Duration, role common.Role) (string, *Payload, error) {
	payload, err := NewPayload(accountId, duration, role)
	if err != nil {
		return "", nil, err
	}

	tokenStr := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := tokenStr.SignedString([]byte(t.secretKey))
	if err != nil {
		return "", nil, err
	}
	return token, payload, nil
}

func (t *JWTMaker) VerifyToken(token string) (*Payload, error) {
	funcKey := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrorInvalidToken
		}
		return []byte(t.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, funcKey)
	if err != nil {
		var ver *jwt.ValidationError
		ok := errors.As(err, &ver)
		if ok && errors.Is(ver.Inner, ErrorExpiredToken) {
			return nil, ErrorExpiredToken
		}
		return nil, ErrorInvalidToken
	}
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, err
	}
	return payload, nil
}
