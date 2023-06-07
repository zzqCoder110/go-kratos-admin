package server

import (
	v1 "go-sim/api/backend/v1"
	"go-sim/internal/conf"
	"go-sim/internal/server/middleware"
	"go-sim/internal/service/menu"
	"go-sim/internal/service/officer"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, cj *conf.Jwt, Officer *officer.OfficerService, Menu *menu.MenuService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		middleware.NewMiddleware(cj),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterOfficerHTTPServer(srv, Officer)
	v1.RegisterMenuHTTPServer(srv, Menu)
	return srv
}
