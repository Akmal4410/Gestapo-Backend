package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Different types of error returend by token
var (
	ErrorExpiredToken error = fmt.Errorf("token is expired")
	ErrorInvalidToken error = fmt.Errorf("token is invalid")
)

// SessionPayload contains the payload data of the session token
type SessionPayload struct {
	Value     string `json:"value"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}

// AccessPayload contains the payload data of the token
type AccessPayload struct {
	UserID    string `json:"user_id"`
	UserName  string `json:"user_name"`
	UserType  string `json:"user_type"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}

// NewSessionPayload creates a new token payload with a specific value and duration
func NewSessionPayload(value, tokenType string, duration time.Duration) *SessionPayload {
	payload := &SessionPayload{
		Value:     value,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	return payload
}

// NewAccessPayload creates a new token payload with a specific username and duration
func NewAccessPayload(userID, userName, userType, tokenType string, duration time.Duration) *AccessPayload {
	payload := &AccessPayload{
		UserID:    userID,
		UserName:  userName,
		UserType:  userType,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	return payload
}

// Valid checks if the token payload is valid or not
func (payload *SessionPayload) Valid() error {
	if time.Now().After(payload.RegisteredClaims.ExpiresAt.Time) {
		return ErrorExpiredToken
	}
	return nil
}

// Valid checks if the token payload is valid or not
func (payload *AccessPayload) Valid() error {
	if time.Now().After(payload.RegisteredClaims.ExpiresAt.Time) {
		return ErrorExpiredToken
	}
	return nil
}
