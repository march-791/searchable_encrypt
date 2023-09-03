package router

import (
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"gonum.org/v1/gonum/mat"
	"gotest/account"
	"gotest/db"
	"gotest/tool"
	"net/http"
	"strings"
)

func Upload(c *gin.Context) {
	//首先认证
	token := c.PostForm("token")
	if token == "" {
		c.String(422, "缺少参数:token")
		return
	}
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
	a := tool.GetCurrentAbPath() + "/file/" + file_info.UserId + "/" + file_info.Path + "/" + file.Filename
	a = strings.Replace(a, "\\", "/", -1)
	filestat := tool.FileState(a)
	fmt.Println("上传文件："+a, filestat)
	err = c.SaveUploadedFile(file, a)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("保存上传文件失败 %s", err.Error()))
		return
	}
	//文件名
	file_info.FileName = file.Filename

	//索引信息
	//file, err = c.FormFile("index")
	//if err != nil {
	//	c.String(http.StatusBadRequest, fmt.Sprintf("获取上传文件失败 %s", err.Error()))
	//	return
	//}
	//File, err := file.Open()
	//tmp := make([]byte, 12800000)
	//_, err = File.Read(tmp)
	//if err != nil && err != io.EOF { //io.EOF代表读取到空文件或文件结尾，这个错误得排除掉
	//	c.String(500, "file read err:", err)
	//	return
	//}
	//index := string(tmp)
	index := c.PostForm("index")
	if index == "" {
		c.String(422, "缺少参数:index")
		return
	}
	var err2 error
	byteindex, _ := hex.DecodeString(index)
	var I []*mat.VecDense
	err2 = tool.Decode(byteindex, &I)
	//if err1 != nil {
	//	c.String(500, "建立索引失败："+err1.Error())
	//	return
	//}
	if err2 != nil {
		c.String(500, "建立索引失败："+err2.Error())
		return
	}
	file_info.I1, _ = I[0].MarshalBinary()
	file_info.I2, _ = I[1].MarshalBinary()

	if filestat == false || db.FileExist(file_info) == false {
		err = db.InsertFile(file_info)
		fmt.Println("插入文件信息")
	} else {
		err = db.UpDateFile(file_info)
		fmt.Println("更新文件信息")
	}
	if err != nil {
		c.String(500, "建立索引失败："+err.Error())
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("上传文件 %s 成功", file.Filename))

}
