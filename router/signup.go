package router

import (
	"github.com/gin-gonic/gin"
	"gotest/db"
)

func Signup(c *gin.Context) {
	info := db.UserInfo{
		UserId:   "",
		Mobile:   "",
		PassWord: "",
	}
	info.UserId = c.PostForm("uid")
	info.PassWord = c.PostForm("password")
	info.Mobile = c.PostForm("mobile")
	//存在缺失的参数
	if info.UserId == "" || info.PassWord == "" || info.Mobile == "" {
		c.String(422, "缺少参数")
		return
	}
	//存入数据库
	err := db.InsertUser(info)
	if err != nil {

		c.String(500, "注册失败", err.Error())
		return
	}
	c.String(200, "Hello, %s", info.UserId)
}
