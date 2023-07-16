package connection

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func GetMySQL(dsn string) (*gorm.DB, error) {
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db, err := conn.DB()
	if err != nil {
		return nil, err
	}
	//设置最大连接数
	db.SetMaxOpenConns(10)
	//设置最大岛空闲连接数
	db.SetMaxIdleConns(5)
	//设置重用连接的最长时间
	db.SetConnMaxLifetime(time.Second * 600)
	return conn, nil
}
