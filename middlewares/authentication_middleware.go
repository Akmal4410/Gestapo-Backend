package middlewares

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/akmal4410/gestapo/helpers"
	"github.com/akmal4410/gestapo/services/token"
)

// Define a custom type for the context key
type contextKey string

const (
	AuthorizationKey        string     = "Authorization"
	AuthorizationTypeBearer string     = "bearer"
	AuthorizationPayloadKey contextKey = "authorization_payload"
)

func AuthenticationMiddleware(tokenMaker token.Maker, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get(AuthorizationKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			helpers.ErrorJson(w, http.StatusUnauthorized, err)
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			helpers.ErrorJson(w, http.StatusUnauthorized, err)
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != AuthorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type: %s", authorizationType)
			helpers.ErrorJson(w, http.StatusUnauthorized, err)
			return
		}

		token := fields[1]

		payload, err := tokenMaker.VerifySessionToken(token)
		if err != nil {
			helpers.ErrorJson(w, http.StatusUnauthorized, err)
			return
		}
		ctx := context.WithValue(r.Context(), AuthorizationPayloadKey, payload)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
