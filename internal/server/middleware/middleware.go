package middleware

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go-sim/internal/conf"
)

func NewMiddleware(jwt *conf.Jwt) http.ServerOption {
	return http.Middleware(
		recovery.Recovery(),
		jwtAuthMiddleware(jwt),
	)
}
