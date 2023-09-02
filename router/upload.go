package router

import (
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"gotest/account"
	"gotest/db"
	"gotest/tool"
	"net/http"
	"strings"
)

func Upload(c *gin.Context) {
	//首先认证
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
	//文件信息
	var file_info db.FileInfo = db.FileInfo{
		UserId:   "",
		Path:     "",
		FileName: "",
		I1:       nil,
		I2:       nil,
	}
	//uid
	file_info.UserId = auth.UserID
	//文件获取路径
	filepath := c.PostForm("path")
	if filepath == "" {
		c.String(422, "缺少参数:path")
		return
	}

	file_info.Path = filepath
	//读取文件
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("获取上传文件失败 %s", err.Error()))
		return
	}

	// 保存文件到本地
	a := tool.GetCurrentAbPath() + "/file/" + "/" + file_info.UserId + "/" + file_info.Path + "/" + file.Filename
	a = strings.Replace(a, "/", "\\", -1)
	fmt.Println("上传文件：" + a)
	err = c.SaveUploadedFile(file, a)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("保存上传文件失败 %s", err.Error()))
		return
	}
	//文件名
	file_info.FileName = file.Filename

	//索引信息
	i1 := c.PostForm("i1")
	i2 := c.PostForm("i2")
	if i1 == "" || i2 == "" {
		c.String(422, "缺少参数i1或i2")
	}
	var err1, err2 error
	file_info.I1, err1 = hex.DecodeString(i1)
	file_info.I2, err2 = hex.DecodeString(i2)
	if err1 != nil || err2 != nil {
		c.String(500, "建立索引失败："+err.Error())
	}
	err = db.InsertFile(file_info)
	if err != nil {
		c.String(500, "建立索引失败："+err.Error())
	}
	c.String(http.StatusOK, fmt.Sprintf("上传文件 %s 成功", file.Filename))

}
