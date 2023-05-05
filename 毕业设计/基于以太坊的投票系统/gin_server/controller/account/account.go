package account

import (
	"BridgeModule/model"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user model.User
	var err error
	//获取数据并绑定的结构体
	err = c.Bind(&user)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"msg": false})
		return
	}
	//
	err = model.Register(&user)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"msg": false})
		return
	}
	//设置session
	session := sessions.Default(c)
	session.Set("id", user.Id)
	session.Set("username", user.Uname)
	session.Set("password", user.Password)
	session.Set("phone", user.Phone)
	session.Options(sessions.Options{MaxAge: 2592000})
	_ = session.Save()
	c.JSON(200, gin.H{"msg": true})
}

func Login(c *gin.Context) {
	var err error
	//获取session数据
	session := sessions.Default(c)
	//获取前端数据
	var user model.User
	err = c.Bind(&user)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"msg": false})
		return
	}
	//比较手机号
	if session.Get("phone").(string) != user.Phone {
		//查询数据库
		data, err := model.GetDataByPhone(user.Phone)
		if err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{"msg": false})
			return
		}
		session.Set("id", data.Id)
		session.Set("username", data.Uname)
		session.Set("password", data.Password)
		session.Set("phone", data.Phone)
		session.Set("ethaddr", data.Ethaddr)
		session.Set("keystore", data.Keystore)
		session.Options(sessions.Options{MaxAge: 2592000})
		session.Save()
	}
	//比较密码
	if user.Password != session.Get("password").(string) {
		c.JSON(200, gin.H{"msg": false})
		return
	}
	c.JSON(200, gin.H{"msg": true})
}
