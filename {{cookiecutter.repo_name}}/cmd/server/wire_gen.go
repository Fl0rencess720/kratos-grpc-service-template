// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"{{cookiecutter.module_name}}/app/{{cookiecutter.repo_name}}/internal/biz"
	"{{cookiecutter.module_name}}/app/{{cookiecutter.repo_name}}/internal/conf"
	"{{cookiecutter.module_name}}/app/{{cookiecutter.repo_name}}/internal/data"
	"{{cookiecutter.module_name}}/app/{{cookiecutter.repo_name}}/internal/server"
	"{{cookiecutter.module_name}}/app/{{cookiecutter.repo_name}}/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	{{cookiecutter.repo_name}}Repo := data.New{{cookiecutter.service_name}}Repo(dataData, logger)
	{{cookiecutter.repo_name}}Usecase := biz.New{{cookiecutter.service_name}}Usecase({{cookiecutter.repo_name}}Repo, logger)
	{{cookiecutter.repo_name}}Service := service.New{{cookiecutter.service_name}}Service({{cookiecutter.repo_name}}Usecase)
	httpServer := server.NewHTTPServer(confServer, {{cookiecutter.repo_name}}Service, logger)
	grpcServer := server.NewGRPCServer(confServer, {{cookiecutter.repo_name}}Service, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
