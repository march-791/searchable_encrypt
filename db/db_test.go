package db

import (
	"fmt"
	"testing"
)

func TestDB_init(t *testing.T) {
	DB_init()
}
func TestInsertUser(t *testing.T) {
	DB_init()
	info := UserInfo{
		UserId:   "test",
		Mobile:   "123456789",
		PassWord: "asdsadsadsdasdas",
	}
	err := InsertUser(info)
	if err != nil {
		return
	}

}
func TestSearchUserByID(t *testing.T) {
	DB_init()
	info := SearchUserByID("test")
	fmt.Println(info)
}
