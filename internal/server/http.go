package server

import (
	v1 "ai-mkt-be/api/filmclip/v1"
	"ai-mkt-be/internal/conf"
	"ai-mkt-be/internal/service"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.FilmclipService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(
				tracing.WithTracerName("filmclip-http"),
				tracing.WithTracerProvider(tracesdk.NewTracerProvider()),
			),
			logging.Server(logger),
			validate.Validator(),
		),
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
	v1.RegisterFilmclipHTTPServer(srv, greeter)
	return srv
}
