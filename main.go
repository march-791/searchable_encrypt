package main

import (
	"github.com/gin-gonic/gin"
	"gotest/db"
	"gotest/log"
	"gotest/router"
)

func main() {
	db.DB_init()
	log.LogInit()
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
	r.POST("/search", func(c *gin.Context) {
		router.Search(c)
	})
	r.POST("/search_t", func(c *gin.Context) {
		router.SearchT(c)
	})
	r.GET("/download", func(c *gin.Context) {
		router.Download(c)
	})
	r.Run() // listen and serve on 0.0.0.0:8080

}
