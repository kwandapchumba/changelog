package middleware

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/kwandapchumba/prioritize/response"
	"github.com/kwandapchumba/prioritize/token"
)

type ContextKey string

const pLoad ContextKey = "payload"

func Authenticator() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			payload, err := parseToken(r)
			if err != nil {
				log.Println(fmt.Errorf("could not parse token cause %s", err.Error()))
				response.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			if payload != nil {

				ctx := context.WithValue(r.Context(), pLoad, payload)

				next.ServeHTTP(w, r.WithContext(ctx))
			}
		}

		return http.HandlerFunc(fn)
	}
}

func parseToken(r *http.Request) (*token.Payload, error) {
	accessToken := r.Header.Get("authorization")

	if accessToken == "" {
		return nil, errors.New("token is empty")
	}

	splitToken := strings.Split(accessToken, "Bearer")

	if len(splitToken) != 2 {
		return nil, errors.New("bearer token is not in proper format")
	}

	accessToken = splitToken[1]

	accessToken = strings.TrimSpace(accessToken)

	payload, err := token.ParseToken(accessToken)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
