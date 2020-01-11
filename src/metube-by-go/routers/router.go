package routers

import (
	"my-blog-by-go/controllers"
	//"log"

	"github.com/gin-gonic/gin"
	//"my-blog-by-go/models"
)

// LoadRouters 初始化router
func LoadRouters(router *gin.Engine) {
	loadRouters(router)
}

func loadRouters(router *gin.Engine) {

	router.GET("/rankinglist", controllers.GetVideos)

	/*
	// 发文
	router.POST("/upload", controllers.UpLoadFile)
	router.POST("/article", controllers.InsertPosts)

	// 点赞
	router.POST("/article/:postid/like", controllers.LikePosts)

	// 查询发文列表
	router.GET("/article", controllers.GetPosts)

	// 查询某个发文详情
	router.GET("/article/:postid", controllers.GetPostById)

	// 查询标签列表
	router.GET("/get-labels", controllers.GetLabels)

	// 查询分类列表
	router.GET("/get-categories", controllers.GetCategoies)

	// 根据label查询文章
	router.GET("/get-posts-by-label/:labelid", controllers.GetPostByLabelId)

	// 根据category查询文章
	router.GET("/get-posts-by-category/:categoryid", controllers.GetPostByCategoryId)
	 */
}
