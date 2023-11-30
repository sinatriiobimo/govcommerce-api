package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/justinas/alice"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
	"net/http"
	"strings"
)

type (
	Skipper func(ctx context.Context, path string) bool

	JWTConfig struct {
		Skipper       Skipper
		SigningKey    []byte
		SigningMethod jwa.SignatureAlgorithm
		ContextKey    interface{}
	}

	JWTUser struct {
		ID int64
	}
)

var (
	DefaultJWTConfig = JWTConfig{
		Skipper:       func(ctx context.Context, path string) bool { return false },
		SigningMethod: jwa.HS256,
		ContextKey:    ContextKeyUser,
	}
)

func JWT(key []byte) alice.Constructor {
	c := DefaultJWTConfig
	c.SigningKey = key
	return JWTWithConfig(c)
}

func JWTWithConfig(config JWTConfig) alice.Constructor {
	if config.Skipper == nil {
		config.Skipper = DefaultJWTConfig.Skipper
	}

	if config.SigningKey == nil {
		panic("jwt middleware requires signing key")
	}

	if config.SigningMethod == "" {
		config.SigningMethod = DefaultJWTConfig.SigningMethod
	}

	if config.ContextKey == "" {
		config.ContextKey = DefaultJWTConfig.ContextKey
	}

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if config.Skipper(r.Context(), r.URL.Path) {
				h.ServeHTTP(w, r)
				return
			}

			token := r.Header.Get("Authorization")
			tokenPrefix := "Bearer "
			if !strings.HasPrefix(token, tokenPrefix) {
				//user token is not mandatory, so we can proceed request without token
				h.ServeHTTP(w, r)
				return
			}
			token = strings.ReplaceAll(token, tokenPrefix, "")

			parsedToken, err := config.parseToken(r.Context(), token)
			if err != nil {
				setError(w, http.StatusUnauthorized, fmt.Errorf("invalid or expired jwt: %v", err))
				return
			}

			r = r.WithContext(context.WithValue(r.Context(), config.ContextKey, parsedToken))

			h.ServeHTTP(w, r)
		})
	}
}

func (config *JWTConfig) parseToken(ctx context.Context, payload string) (interface{}, error) {
	token, err := jwt.Parse(
		[]byte(payload),
		jwt.WithValidate(true),
		jwt.WithVerify(config.SigningMethod, config.SigningKey),
	)

	if err != nil {
		return nil, err
	}

	user, err := token.AsMap(ctx)
	if err != nil {
		return nil, err
	}
	userId, ok := user["id"].(float64)
	if !ok {
		return nil, errors.New("cannot cast user id from token")
	}
	return JWTUser{
		ID: int64(userId),
	}, err
}
