package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"middlewares/internal/config"
	"middlewares/internal/middleware"
)

type ServiceContext struct {
	Config         config.Config
	UserMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserMiddleware: middleware.NewUserMiddleware().Handle,
	}
}
