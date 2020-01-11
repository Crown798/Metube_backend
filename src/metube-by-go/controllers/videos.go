package controllers

import (
	"my-blog-by-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetPosts 获取所有的
func GetVideos(c *gin.Context) {
	videos := models.GetVideos()
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":  videos,
		"code": 0,
	})
}

/*
//GetPostByLabelId 根据label id获取post
func GetPostByLabelId(c *gin.Context) {
	labelid := c.Param("labelid")
	if labelid == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 1,
			"msg":    "labelid 不能为空",
		})
	}
	labelId, err := strconv.ParseInt(labelid, 10, 64)
	postIDs, err := models.GetPostsByLabelID(labelId)
	if err != nil {
		log.Println(labelId, err)
		c.JSON(http.StatusNotFound, gin.H{
			"status": 1,
			"msg":    "GetPostsByLabelID error, err:" + err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   models.GetPostByIDs(postIDs),
	})
}

//GetPosts 获取所有的文章
func InsertPosts(c *gin.Context) {

	requestInfo := models.InsertPostReq{}
	err := c.BindJSON(&requestInfo)
	if err != nil {
		fmt.Println("Parse body failed. Error:", err)
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "params 格式错误",
		})
		return
	}

	// 检查文章是否已经存在
	postInfo := models.GetPostByTitle(requestInfo.Title)
	if postInfo != nil && postInfo.Id > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": 2,
			"msg":    "title 已存在",
		})
		return
	}

	id, err := models.InsertPostByTransaction(requestInfo)
	if err != nil {
		log.Println("InsertPostByTransaction failed. Error:", err)
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "transaction error",
		})
		return
	}

	posts := models.GetPosts()
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"data":    posts,
		"post_id": id,
		"msg":     "success",
	})
}
*/
