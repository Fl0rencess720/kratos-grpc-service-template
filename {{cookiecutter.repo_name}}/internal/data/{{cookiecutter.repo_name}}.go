package data

import (
	"{{cookiecutter.module_name}}/app/{{cookiecutter.repo_name}}/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type {{cookiecutter.repo_name}}Repo struct {
	data *Data
	log  *log.Helper
}

func New{{cookiecutter.service_name}}Repo(data *Data, logger log.Logger) biz.{{cookiecutter.service_name}}Repo {
	return &{{cookiecutter.repo_name}}Repo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

