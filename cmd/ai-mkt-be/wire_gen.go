// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"ai-mkt-be/internal/biz"
	"ai-mkt-be/internal/conf"
	"ai-mkt-be/internal/data"
	"ai-mkt-be/internal/server"
	"ai-mkt-be/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	agentGraph := biz.NewAgentGraph(logger)
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	filmclipService := service.NewFilmclipService(logger, agentGraph, greeterUsecase)
	grpcServer := server.NewGRPCServer(confServer, filmclipService, logger)
	httpServer := server.NewHTTPServer(confServer, filmclipService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
