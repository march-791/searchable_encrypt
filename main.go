package main

import (
	"github.com/gin-gonic/gin"
	"gotest/db"
	"gotest/router"
)

func main() {
	db.DB_init()
	r := gin.Default()
	//注册
	r.POST("/signup", func(c *gin.Context) {
		router.Signup(c)
	})
	r.POST("/login", func(c *gin.Context) {
		router.Login(c)
	})
	r.POST("/auth", func(c *gin.Context) {
		router.Auth(c)
	})
	r.POST("/upload", func(c *gin.Context) {
		router.Upload(c)
	})
	r.Run() // listen and serve on 0.0.0.0:8080

}
