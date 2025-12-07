package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secret string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {

		return nil, fmt.Errorf("invalid key size. must be atleast %d characters", minSecretKeySize)
	}

	return &JWTMaker{secretKey}, nil

}

// create a token
func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	jwtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return jwtoken.SignedString([]byte(maker.secret))
}

// verify a token
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {

	keyfunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken

		}
		return []byte(maker.secret), nil

	}

	jwttoken, err := jwt.ParseWithClaims(token, &Payload{}, keyfunc)
	if err != nil {

		if errors.Is(err, jwt.ErrTokenExpired) {

			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwttoken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil

}
