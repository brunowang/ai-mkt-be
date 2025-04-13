//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"ai-mkt-be/internal/aigc"
	"ai-mkt-be/internal/biz"
	"ai-mkt-be/internal/conf"
	"ai-mkt-be/internal/data"
	"ai-mkt-be/internal/lib"
	"ai-mkt-be/internal/server"
	"ai-mkt-be/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		lib.ProviderSet,
		aigc.ProviderSet,
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		newApp))
}
