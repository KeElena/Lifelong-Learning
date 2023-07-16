package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user/model"

	"user/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Conn   model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config: c,
		Conn:   model.NewUsersModel(sqlConn, c.CacheRedis),
	}
}
