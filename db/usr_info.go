package db

import (
	"fmt"
)

func InsertUser(info UserInfo) error {
	if err := Database.Create(&info).Error; err != nil {
		fmt.Println("插入失败err" + err.Error())
		return err
	}
	return nil
}
func SearchUserByID(id string) UserInfo {
	var info UserInfo
	Database.First(&info, "user_id = ?", id)
	return info
}
