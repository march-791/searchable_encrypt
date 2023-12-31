package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gotest/account"
	"gotest/db"
)

func Login(c *gin.Context) {
	info := db.UserInfo{
		UserId:   "",
		Mobile:   "",
		PassWord: "",
	}
	info.UserId = c.PostForm("uid")
	info.PassWord = c.PostForm("password")
	fmt.Println(info.UserId + "     " + info.PassWord)
	//存在缺失的参数
	if info.UserId == "" || info.PassWord == "" {
		c.String(422, "缺少参数")
		return
	}
	//登陆获取token
	token, err := account.Login(info.UserId, info.PassWord)
	if err != nil {
		c.String(422, "错误："+err.Error())
		return
	}
	c.String(200, token)
}
