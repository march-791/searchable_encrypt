package router

import (
	"github.com/gin-gonic/gin"
	"gotest/account"
)

func Auth(c *gin.Context) {
	token := c.PostForm("token")
	auth, err := account.Authority(token)
	if err != nil {
		c.String(401, "认证失败："+err.Error())
		return
	}
	if auth.Auth == false {
		c.String(401, auth.UserID+"认证失败")
		return
	}
	c.String(200, auth.UserID+"认证成功")
}
