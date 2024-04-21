package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/akmal4410/gestapo/pkg/helpers"
	"github.com/akmal4410/gestapo/pkg/helpers/logger"
	"github.com/akmal4410/gestapo/pkg/helpers/token"
	"github.com/akmal4410/gestapo/pkg/utils"
)

func AuthMiddleware(tokenMaker token.Maker, log logger.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get(utils.AuthorizationKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			log.LogError("Error", err)
			helpers.ErrorJson(w, http.StatusUnauthorized, err.Error())
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			log.LogError("Error", err)
			helpers.ErrorJson(w, http.StatusUnauthorized, err.Error())
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != utils.AuthorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type: %s", authorizationType)
			log.LogError("Error", err)
			helpers.ErrorJson(w, http.StatusUnauthorized, err.Error())
			return
		}

		token := fields[1]

		payload, err := tokenMaker.VerifySessionToken(token)
		if err != nil {
			log.LogError("Error", err)
			helpers.ErrorJson(w, http.StatusUnauthorized, err.Error())
			return
		}
		ctx := context.WithValue(r.Context(), utils.AuthorizationPayloadKey, payload)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func SsoMiddleware(log logger.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get(utils.AuthorizationKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			log.LogError("Error", err)
			helpers.ErrorJson(w, http.StatusUnauthorized, err.Error())
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			log.LogError("Error", err)
			helpers.ErrorJson(w, http.StatusUnauthorized, err.Error())
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != utils.AuthorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type: %s", authorizationType)
			log.LogError("Error", err)
			helpers.ErrorJson(w, http.StatusUnauthorized, err.Error())
			return
		}

		token := fields[1]

		ctx := context.WithValue(r.Context(), utils.AuthorizationPayloadKey, token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
