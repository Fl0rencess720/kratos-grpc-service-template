package biz

import (
	"context"

	v1 "{{cookiecutter.module_name}}/api/gateway/{{cookiecutter.repo_name}}/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// {{cookiecutter.service_name}} is a {{cookiecutter.service_name}} model.
type {{cookiecutter.service_name}} struct {
	
}

// {{cookiecutter.service_name}}Repo is a Greater repo.
type {{cookiecutter.service_name}}Repo interface {

}

// {{cookiecutter.service_name}}Usecase is a {{cookiecutter.service_name}} usecase.
type {{cookiecutter.service_name}}Usecase struct {
	repo {{cookiecutter.service_name}}Repo
	log  *log.Helper
}

// New{{cookiecutter.service_name}}Usecase new a {{cookiecutter.service_name}} usecase.
func New{{cookiecutter.service_name}}Usecase(repo {{cookiecutter.service_name}}Repo, logger log.Logger) *{{cookiecutter.service_name}}Usecase {
	return &{{cookiecutter.service_name}}Usecase{repo: repo, log: log.NewHelper(logger)}
}
