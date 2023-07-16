package svc

import (
	"gorm.io/gorm"
	"user/connection"
	"user/internal/config"
	"user/internal/model"
)

type ServiceContext struct {
	Config  config.Config
	MysqlDB *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	MysqlDB, err := connection.GetMySQL(c.Mysql.DataSource)
	if err != nil {
		panic("数据库连接失败!")
	}
	//同步结构体到数据库
	_ = MysqlDB.AutoMigrate(&model.User{})
	return &ServiceContext{
		Config: c,
		//返回数据库连接
		MysqlDB: MysqlDB,
	}
}
