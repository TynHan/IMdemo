package models

import (
	"ginchat/utils"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Salt          string
	Phone         string `valid:"matches(^13\\d{9}$)"`
	Email         string `valid:"email"`
	Identity      string
	ClientIP      string
	ClinetPort    string
	LoginTime     uint64
	HeartbeatTime uint64
	LogoutTime    uint64
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	db := utils.GetDB()
	// res := make([]*UserBasic, 10)
	res := []*UserBasic{}
	db.Find(&res)
	return res
}

func GetUserByName(name string) []*UserBasic {
	db := utils.GetDB()
	res := []*UserBasic{}
	db.Model(&UserBasic{}).Where("name = ?", name).Find(&res)
	return res
}

func GetUserByNameAndPwd(name, password string) []*UserBasic {
	db := utils.GetDB()
	res := []*UserBasic{}
	db.Model(&UserBasic{}).Where("name = ? and password = ?", name, password).Find(&res)
	return res
}

func GetUserByPhone(phone string) []*UserBasic {
	db := utils.GetDB()
	res := []*UserBasic{}
	db.Model(&UserBasic{}).Where("phone = ?", phone).Find(&res)
	return res
}

func GetUserByEmail(email string) []*UserBasic {
	db := utils.GetDB()
	res := []*UserBasic{}
	db.Model(&UserBasic{}).Where("email = ?", email).Find(&res)
	return res
}

func CreateUser(user *UserBasic) {
	db := utils.GetDB()
	db.Create(user)
}

func DeleteUser(user *UserBasic) {
	db := utils.GetDB()
	db.Delete(user)
}

func UpdateUser(user *UserBasic) {
	db := utils.GetDB()
	db.Model(user).Select("Name", "Email", "Phone").Updates(user)
	// db.Save(user)
}

func UpdateUserTocken(user *UserBasic) {
	db := utils.GetDB()
	db.Model(user).Select("Identity").Updates(user)
	// db.Save(user)
}
