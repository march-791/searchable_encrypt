package db

import "fmt"

func InsertFile(info FileInfo) error {
	if err := Database.Create(&info).Error; err != nil {
		fmt.Println("插入失败err" + err.Error())
		return err
	}
	return nil
}
