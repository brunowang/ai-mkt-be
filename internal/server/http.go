package server

import (
	v1 "ai-mkt-be/api/filmclip/v1"
	"ai-mkt-be/internal/conf"
	"ai-mkt-be/internal/service"
	"encoding/json"
	kjson "github.com/go-kratos/kratos/v2/encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	nethttp "net/http"
	"time"
)

type HTTPResponse struct {
	Code    int     `json:"code"`
	Mesg    string  `json:"mesg"`
	Time    float64 `json:"time"`
	Data    any     `json:"data,omitempty"`
	TraceId string  `json:"trace_id"`
}

func init() {
	kjson.MarshalOptions.UseProtoNames = true
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, filmclip *service.FilmclipService, logger log.Logger) *http.Server {
	nethttp.HandleFunc("/", NotFoundHandler)
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(
				tracing.WithTracerName("filmclip-http"),
				tracing.WithTracerProvider(tracesdk.NewTracerProvider()),
			),
			logging.Server(logger),
			validate.Validator(),
			MiddlewareCors(),
		),
		http.ResponseEncoder(func(w nethttp.ResponseWriter, req *nethttp.Request, rsp any) error {
			var traceId string
			if r, ok := rsp.(*service.RspContext); ok && r != nil {
				traceId = tracing.TraceID()(r.Ctx).(string)
			}
			bs, err := json.Marshal(&HTTPResponse{
				Code:    0,
				Mesg:    "处理成功",
				Time:    float64(time.Now().UnixNano()) / 1e9,
				Data:    rsp,
				TraceId: traceId,
			})
			if err != nil {
				return err
			}
			w.Header().Set("Content-Type", "application/json")
			_, err = w.Write(bs)
			return err
		}),
		http.ErrorEncoder(func(w nethttp.ResponseWriter, req *nethttp.Request, err error) {
			var traceId string
			if e, ok := err.(*service.ErrContext); ok && e != nil {
				traceId = tracing.TraceID()(e.Ctx).(string)
			}
			code, mesg := 500, "处理失败"
			if err != nil {
				mesg = err.Error()
			}
			bs, _ := json.Marshal(&HTTPResponse{
				Code:    code,
				Mesg:    mesg,
				Time:    float64(time.Now().UnixNano()) / 1e9,
				TraceId: traceId,
			})
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write(bs)
		}),
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
	v1.RegisterFilmclipHTTPServer(srv, filmclip)
	return srv
}
