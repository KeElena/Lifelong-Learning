package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"rpc-common/userclient"
	"rpc-common/userscore"
	"userapi/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserRpc   userclient.User
	UserScore userscore.UserScore
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserRpc:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		UserScore: userscore.NewUserScore(zrpc.MustNewClient(c.UserScore)),
	}
}
