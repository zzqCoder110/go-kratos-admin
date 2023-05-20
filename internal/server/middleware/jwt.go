package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	"go-sim/internal/conf"
)

func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/backend.v1.Officer/Login"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

func jwtAuthMiddleware(jwtConf *conf.Jwt) middleware.Middleware {
	return selector.Server(
		jwt.Server(func(token *jwtV4.Token) (interface{}, error) {
			return []byte(jwtConf.Key), nil
		}, jwt.WithSigningMethod(jwtV4.SigningMethodHS256), jwt.WithClaims(func() jwtV4.Claims {
			return &jwtV4.MapClaims{}
		})),
	).
		Match(NewWhiteListMatcher()).
		Build()
}
