package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// ORM gorm引擎的实例
var ORM *gorm.DB

func Init() {
	var err error

	//连接数据库，为方便教学，数据库使用sqlite
	db, err := gorm.Open("mysql", "root:12345678@tcp(123.56.96.92:3306)/metube?charset=utf8&parseTime=true")
	if err != nil {
		panic("连接数据库失败")
	}

	ORM = db
	// 打印sql详情
	db.LogMode(true)
}
