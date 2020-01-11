package models

import (
	"my-blog-by-go/database"
)

func AutoMigrate() {

	// 自动迁移数据库格式
	database.ORM.AutoMigrate(&MVideo{}, &MUser{})
}
