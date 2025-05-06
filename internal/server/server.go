package server

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	nethttp "net/http"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer)

// MiddlewareCors 设置跨域请求头
func MiddlewareCors() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if ts, ok := transport.FromServerContext(ctx); ok {
				if ht, ok := ts.(http.Transporter); ok {
					ht.ReplyHeader().Set("Access-Control-Allow-Origin", "*")
					ht.ReplyHeader().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,PUT,PATCH,DELETE")
					ht.ReplyHeader().Set("Access-Control-Allow-Credentials", "true")
					ht.ReplyHeader().Set("Access-Control-Allow-Headers", "Content-Type,"+
						"X-Requested-With,Access-Control-Allow-Credentials,User-Agent,Content-Length,Authorization")
				}
			}
			return handler(ctx, req)
		}
	}
}

func NotFoundHandler(res nethttp.ResponseWriter, req *nethttp.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,PUT,PATCH,DELETE")
	res.Header().Set("Access-Control-Allow-Credentials", "true")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type,"+
		"X-Requested-With,Access-Control-Allow-Credentials,User-Agent,Content-Length,Authorization")

	log.Info("NotFoundHandler")

	errCode := errors.NotFound("page not found", "page not found")
	buffer, _ := json.Marshal(errCode)
	//res.WriteHeader(400)
	_, _ = res.Write(buffer)
}
