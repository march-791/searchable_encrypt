package db

type UserInfo struct {
	UserId   string `gorm:"column:user_id"`
	Mobile   string `gorm:"column:mobile"`
	PassWord string `gorm:"column:password"`
}

func (p UserInfo) TableName() string {
	return "user_info"
}
