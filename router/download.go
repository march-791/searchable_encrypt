package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gotest/account"
	"gotest/tool"
	"os"
	"strings"
)

func Download(c *gin.Context) {
	token := c.Query("token")
	claims, err := account.ParseToken2(token)
	if err != nil {
		c.String(500, "解析token失败"+err.Error())
	}
	fmt.Println(claims)
	path := claims.Path
	name := strings.Trim(claims.FileName, " ")
	// path

	filePath := strings.Replace(path, "\\", "/", -1) + name
	abspath := tool.GetCurrentAbPath() + "/file/" + filePath
	_, _ = os.Open(abspath)
	fmt.Println(abspath)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+name)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(abspath)
}
