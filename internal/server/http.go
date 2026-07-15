package server

import (
	user "bin-personal-book/api/user/v1"
	"bin-personal-book/internal/conf"
	"bin-personal-book/internal/service"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtV5 "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/handlers"
)

// 过滤白名单
func NewWhiteListMatcher() selector.MatchFunc {
	var whiteList = map[string]struct{}{
		"/api.user.v1.Greeter/Login":    {},
		"/api.user.v1.Greeter/Register": {},
	}

	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Bootstrap, greeter *service.MainService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),

			// 日志
			logging.Server(logger),

			// jwt 校验
			selector.Server(
				jwt.Server(func(token *jwtV5.Token) (interface{}, error) {
					return []byte(c.Data.Jwt.Secret), nil
				}),
			).
				Match(NewWhiteListMatcher()).
				Build(),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Cookie"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}

	if c.Server.Http.Network != "" {
		opts = append(opts, http.Network(c.Server.Http.Network))
	}

	if c.Server.Http.Addr != "" {
		opts = append(opts, http.Address(c.Server.Http.Addr))
	}

	if c.Server.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Server.Http.Timeout.AsDuration()))
	}

	srv := http.NewServer(append(opts,
		http.ResponseEncoder(ResponseEncoder),
		http.ErrorEncoder(ErrorEncoder),
	)...)

	user.RegisterGreeterHTTPServer(srv, greeter)

	return srv
}
