package model

import (
	"fmt"
	"github.com/ylinyang/vue-demo-go/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var Db *gorm.DB

func InitDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser, utils.DbPassWord, utils.DbHost, utils.DbPort, utils.DbName)
	Db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tb_",
			SingularTable: true,
		},
	})
	// 通过type字段自动创建数据库
	if err := Db.AutoMigrate(&User{}); err != nil {
		log.Panicln("迁移数据库失败", err)
	}
}
