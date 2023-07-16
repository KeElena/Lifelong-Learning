package curd

import (
	"gorm.io/gorm"
	"user/internal/model"
)

func Insert(mysqlEngin *gorm.DB, id int64) error {
	return mysqlEngin.Create(&model.User{Id: id, Name: "demo", Gender: "武装直升机"}).Error
}

func Query(mysqlEngin *gorm.DB, id int64) (*model.User, error) {
	var user *model.User
	err := mysqlEngin.Where("id=?", id).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
