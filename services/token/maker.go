package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	// CreateSessionToken create a session token for specific value and duration
	CreateSessionToken(value, tokenType string, duration time.Duration) (string, error)

	// VerifySessionToken  checks if session token is valid or not
	VerifySessionToken(token string) (*SessionPayload, error)

	// CreateAccessToken create a access token for specific userName and duration
	CreateAccessToken(userName string, duration time.Duration) (string, error)

	// VerifyAccessToken checks if access token is valid or not
	VerifyAccessToken(token string) (*AccessPayload, error)
}
