package db

import "fmt"

func InsertFile(info FileInfo) error {
	if err := Database.Create(&info).Error; err != nil {
		fmt.Println("插入失败err" + err.Error())
		return err
	}
	return nil
}
func UpDateFile(info FileInfo) error {
	err := Database.Model(&info).Where("user_id = ? AND path= ? AND file_name =?", info.UserId, info.Path, info.FileName).Update("i1", info.I1).Error
	if err != nil {
		return err
	}
	err = Database.Model(&info).Where("user_id = ? AND path= ? AND file_name =?", info.UserId, info.Path, info.FileName).Update("i2", info.I2).Error
	if err != nil {
		return err
	}
	return nil

}
func FileExist(info FileInfo) bool {
	re := Database.First(&info, "user_id = ? AND path= ? AND file_name =?", info.UserId, info.Path, info.FileName)
	if re.RowsAffected == 0 {
		return false
	}
	return true
}
func PathExist(userId, path string) bool {
	var info FileInfo
	re := Database.First(&info, "user_id = ? AND path= ? ", userId, path)
	if re.RowsAffected == 0 {
		return false
	}
	return true
}
func FindFileInfoByPath(userId, path string) ([]*FileInfo, error) {
	var fileInfo []*FileInfo
	if err := Database.Model(&FileInfo{}).Where("user_id = ? AND path = ?", userId, path).Find(&fileInfo).Error; err != nil {
		return nil, err
	}
	return fileInfo, nil
}
