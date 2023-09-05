package router

import (
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"gonum.org/v1/gonum/mat"
	"gotest/account"
	"gotest/algorithm/search"
	"gotest/db"
	"gotest/tool"
	"strconv"
)

func Search(c *gin.Context) {
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
	//path
	path := c.PostForm("path")

	file_path := auth.UserID + "\\" + path + "\\"
	trapdoorString := c.PostForm("trapdoor")
	if db.PathExist(auth.UserID, path) == false {
		c.String(422, "路径不存在")
		return
	}
	fileInfos, err := db.FindFileInfoByPath(auth.UserID, path)
	if err != nil {
		c.String(500, "查找失败", err.Error())
		return
	}
	indexAndID := make([]*search.IndexAndID, len(fileInfos))
	for i, fileInfo := range fileInfos {
		I1 := make([]*mat.VecDense, 2)
		b1 := new(mat.VecDense)
		b2 := new(mat.VecDense)
		err = b1.UnmarshalBinary(fileInfo.I1)
		if err != nil {
			c.String(500, "处理陷门出错")
			return
		}
		err = b2.UnmarshalBinary(fileInfo.I2)
		if err != nil {
			c.String(500, "处理陷门出错")
			return
		}
		I1[0] = b1
		I1[1] = b2
		file_token, err := account.GenToken2(file_path, fileInfo.FileName)
		if err != nil {
			c.String(500, "Token生成失败")
			return
		}
		url := fmt.Sprintf("%s%s%s%s%s", "http://", "123.56.185.106:",
			"8080", "/download?token=", file_token)
		indexAndID[i] = &search.IndexAndID{
			Id:       strconv.Itoa(i),
			Index:    I1,
			Date:     url,
			FileName: fileInfo.FileName,
		}
	}
	byteTrap, err := hex.DecodeString(trapdoorString)
	var Tp []*mat.VecDense
	err = tool.Decode(byteTrap, &Tp)
	results := search.Search(indexAndID, Tp)
	excptNum, err := strconv.Atoi(c.PostForm("exceptnum"))
	if err != nil {
		c.String(422, "参数错误：exceptnum")
		return
	}

	if len(results) > excptNum && excptNum > 0 {
		results = results[0:excptNum]
	}
	resultBytes, err := tool.Encode(results)
	if err != nil {
		c.String(500, "获取结果失败")
		return
	}
	resultStr := hex.EncodeToString(resultBytes)
	c.String(200, resultStr)
}
