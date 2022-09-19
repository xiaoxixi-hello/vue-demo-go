package model

import (
	"github.com/ylinyang/vue-demo-go/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(80);not null" json:"Password"`
}

// CheckUser 查询用户是否存在
func CheckUser(name string) int {
	var u User
	tx := Db.Where("username = ?", name).First(&u)
	if tx.Error == nil {
		return utils.ErrorUserExist
	}
	return utils.SUCCESS
}

// CreateUser 新增用户
func CreateUser(u *User) int {
	u.Password = hashPassword(u.Password)
	if Db.Create(u).Error != nil {
		log.Fatalln("创建写入数据失败", Db.Create(u).Error)
		return utils.ErrorMysql
	}
	return utils.SUCCESS
}

// GetUsers 查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var u []User
	if Db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&u).Error != nil {
		return nil
	}
	return u
}

// 用户密码加密
func hashPassword(password string) string {
	bytes := []byte(password)
	fromPassword, err := bcrypt.GenerateFromPassword(bytes, 4)
	if err != nil {
		log.Fatalln("密码加密失败", err)
	}
	return string(fromPassword)
}

// CheckLogin 登录验证
func CheckLogin(u, p string) int {
	var user User
	Db.Where("username = ?", u).First(&user)
	if user.ID == 0 {
		return utils.ErrorUserNoExist
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p)); err != nil {
		return utils.ErrorUserPassWord
	}
	return utils.SUCCESS
}
