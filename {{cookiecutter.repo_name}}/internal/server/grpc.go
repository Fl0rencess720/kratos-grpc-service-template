package server

import (
	v1 "{{cookiecutter.module_name}}/api/gateway/{{cookiecutter.repo_name}}/v1"
	"{{cookiecutter.module_name}}/app/{{cookiecutter.repo_name}}/internal/conf"
	"{{cookiecutter.module_name}}/app/{{cookiecutter.repo_name}}/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, {{cookiecutter.repo_name}} *service.{{cookiecutter.service_name}}Service, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			ratelimit.Server(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.Register{{cookiecutter.service_name}}Server(srv, {{cookiecutter.repo_name}})
	return srv
}
