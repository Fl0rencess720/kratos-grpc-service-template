package service

import (
	"context"

	v1 "{{cookiecutter.module_name}}/api/gateway/{{cookiecutter.repo_name}}/v1"
	"{{cookiecutter.module_name}}/app/{{cookiecutter.repo_name}}/internal/biz"
)

type {{cookiecutter.service_name}}Service struct {
	v1.Unimplemented{{cookiecutter.service_name}}Server

	uc *biz.{{cookiecutter.service_name}}Usecase
}

func New{{cookiecutter.service_name}}Service(uc *biz.{{cookiecutter.service_name}}Usecase) *{{cookiecutter.service_name}}Service {
	return &{{cookiecutter.service_name}}Service{uc: uc}
}


