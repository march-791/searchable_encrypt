package db

type UserInfo struct {
	UserId   string `gorm:"column:user_id"`
	Mobile   string `gorm:"column:mobile"`
	PassWord string `gorm:"column:password"`
}

func (p UserInfo) TableName() string {
	return "user_info"
}

type FileInfo struct {
	UserId   string `gorm:"column:user_id"`
	Path     string `gorm:"column:path"`
	FileName string `gorm:"column:file_name"`
	I1       []byte `gorm:"column:i1"`
	I2       []byte `gorm:"column:i2"`
}

func (p FileInfo) TableName() string {
	return "file_info"
}

type SearchFileInfo struct {
	Id string `gorm:"column:id"`
	I1 []byte `gorm:"column:i1"`
	I2 []byte `gorm:"column:i2"`
}

func (p SearchFileInfo) TableName() string {
	return "file_info"
}
