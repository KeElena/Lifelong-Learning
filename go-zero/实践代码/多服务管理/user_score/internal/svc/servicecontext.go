package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user_score/internal/config"
	"user_score/internal/model"
)

type ServiceContext struct {
	Config config.Config
	Conn   model.UserScoreModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		Conn:   model.NewUserScoreModel(sqlConn, c.CacheRedis),
	}
}
