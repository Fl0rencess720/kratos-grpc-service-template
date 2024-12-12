//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"{{cookiecutter.module_name}}/app/{{cookiecutter.repo_name}}/internal/biz"
	"{{cookiecutter.module_name}}/app/{{cookiecutter.repo_name}}/internal/conf"
	"{{cookiecutter.module_name}}/app/{{cookiecutter.repo_name}}/internal/data"
	"{{cookiecutter.module_name}}/app/{{cookiecutter.repo_name}}/internal/server"
	"{{cookiecutter.module_name}}/app/{{cookiecutter.repo_name}}/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Registry, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
