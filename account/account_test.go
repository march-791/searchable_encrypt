package account

import (
	"fmt"
	"gotest/db"
	"testing"
)

func TestAccount(t *testing.T) {
	//用户信息数据结构
	db.DB_init()
	var a db.UserInfo = db.UserInfo{
		UserId:   "01",
		Mobile:   "1851000",
		PassWord: "asdddwqeeeeee",
	}
	//注册
	flag, err := Register(a)
	if flag == 0 || err != nil {
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("注册失败")
		return
	}
	//登陆，返回token
	token, err := Login("01", "asdddwqeeeeee")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("token:")
	fmt.Println(token)
	//parse token
	parse, err := ParseToken(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("parse token:")
	fmt.Println(parse)

	//验证签名
	auth, err := Authority(token)
	if err != nil {
		fmt.Println(err)
	}
	if auth.Auth == true {
		fmt.Printf("%s认证成功", auth.UserID)
	}
}
