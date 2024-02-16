package token

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

const minSecretKeySize = 32

// JWTMaker is JSON Wed Token Maker
type JWTMaker struct {
	secretKey string
}

// CreateSessionToken implements Maker.
func (maker *JWTMaker) CreateSessionToken(value string, duration time.Duration) (string, error) {
	mySigningKey := []byte(maker.secretKey)
	payload := NewSessionPayload(value, duration)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString(mySigningKey)
}

// VerifySessionToken implements Maker.
func (maker *JWTMaker) VerifySessionToken(token string) (*SessionPayload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrorInvalidToken
		}
		return []byte(maker.secretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &SessionPayload{}, keyFunc)
	if err != nil {
		if strings.Contains(err.Error(), "token is expired") {
			return nil, ErrorExpiredToken
		}
		return nil, ErrorInvalidToken
	}
	if !jwtToken.Valid {
		return nil, ErrorInvalidToken
	}
	payload, ok := jwtToken.Claims.(*SessionPayload)
	if !ok {
		return nil, ErrorInvalidToken
	}
	return payload, nil
}

// NewJWTMaker creates a new JWTMaker
func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be atleast %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

// CreateAccessToken create a token for specific userName and duration
func (maker *JWTMaker) CreateAccessToken(userName string, duration time.Duration) (string, error) {
	mySigningKey := []byte(maker.secretKey)
	payload := NewAccessPayload(userName, duration)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString(mySigningKey)

}

// VerifyAccessToken checks if token is valid or not
func (maker *JWTMaker) VerifyAccessToken(token string) (*AccessPayload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrorInvalidToken
		}
		return []byte(maker.secretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &AccessPayload{}, keyFunc)
	if err != nil {
		if strings.Contains(err.Error(), "token is expired") {
			return nil, ErrorExpiredToken
		}
		return nil, ErrorInvalidToken
	}
	if !jwtToken.Valid {
		return nil, ErrorInvalidToken
	}
	payload, ok := jwtToken.Claims.(*AccessPayload)
	if !ok {
		return nil, ErrorInvalidToken
	}
	return payload, nil
}
