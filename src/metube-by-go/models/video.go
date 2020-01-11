package models

import (
	"log"
	db "my-blog-by-go/database"
	"time"
)

type MVideo struct {
	Id         int  `gorm:"column:id;primary_key;AUTO_INCREMENT;not null;" json:"id"`
	CreatedAt  time.Time `gorm:"column:created_at;" json:"-"`
	UpdatedAt  time.Time `gorm:"column:updated_at;" json:"updated_at"`
	DeletedAt  time.Time `gorm:"column:deleted_at;" json:"-"`
	Title      string `gorm:"column:title" json:"title"`
	Info       string `gorm:"column:info" json:"info"`
	Url    	   string `gorm:"column:url" json:"-"`
	Avatar     string `gorm:"column:avatar" json:"avatar"`
	Views      int  `gorm:"column:views" json:"views"`
	Owner      int `gorm:"column:owner" json:"owner"`
	Typename   string `gorm:"column:typename" json:"typename"`
}

type MUser struct {
	Id int `gorm:"column:id;primary_key;AUTO_INCREMENT;not null;" json:"id"`
	Nickname string `gorm:"column:nickname" json:"nickname"`
}

func (p *MVideo) TableName() string {
	return "videos"
}

func (p *MUser) TableName() string {
	return "users"
}

type result struct {
	Id         int  `json:"id"`
	UpdatedAt  time.Time `json:"updated_at"`
	Title      string `json:"title"`
	Info       string `json:"info"`
	Avatar     string `json:"avatar"`
	Views      int  `json:"views"`
	Owner      string `json:"owner"`
	Typename   string `json:"typename"`
}

func GetVideos() []result {
	//var video []*MVideo
	//err := db.ORM.Order("views").Find(&video).Error
	//if err != nil {
	//	log.Println("ERROR:", err)
	//	return nil
	//}
	//
	//length := len(video)
	//for i := 0; i < length/2; i++ {
	//	temp := video[length-1-i]
	//	video[length-1-i] = video[i]
	//	video[i] = temp
	//}
	//return video
	var results []result
	err := db.ORM.Table("videos").Select("videos.id as id, videos.updated_at as updated_at, videos.title as title, videos.info as info, videos.avatar as avatar, videos.views as views, users.nickname as owner, videos.typename as typename").Joins("left join users on videos.owner = users.id").Order("views").Scan(&results).Error
	if err != nil {
		log.Println("ERROR:", err)
		return nil
	}
	length := len(results)
	for i := 0; i < length/2; i++ {
		temp := results[length-1-i]
		results[length-1-i] = results[i]
		results[i] = temp
	}
	return results
}
