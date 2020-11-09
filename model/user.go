package model

type User struct {
	Id int `gorm:"primary_key;auto_increment" json:"id"`
	NickName string `gorm:"size:64;column:nickname;default:''" json:"nickname"`
	Password  string `gorm:"size:256;column:password;default:''" json:"password"`
	Status  string `gorm:"size:256;column:status;default:'1'" json:"status"`
	Avatar  string `gorm:"size:256;column:avatar;default:''" json:"avatar"`
}

func (User) TableName() string {
	return "users"
}
