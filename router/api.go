package router

import (
	"github.com/LINK/controller"
	"github.com/gin-gonic/gin"
)

func apiRouter(api *gin.RouterGroup) {

	// Company系
	api.GET("/companies", controller.GetAllCompanyCards)

	// Category系
	api.GET("/categories", controller.GetCategories)

	// Content系
	api.GET("/categories/contents", controller.GetContentsFilteredByCategory)
	api.GET("/contents/", controller.GetAllContentCards)
	api.GET("/contents/sort/", controller.GetSortedContentCards)
	api.GET("/content/:content_id", controller.GetContentById)

	// Event系
	api.GET("/events", controller.GetAllEventCards)
	api.GET("/event/:event_id", controller.GetEventById)
	//api.GET("/events", controller.GetAllEvents)

	// Homeで必要となるデータをまとめて取得
	api.GET("/home_content", controller.GetHomeContent)

	// Article系
	api.GET("/content/:content_id/articles/", controller.GetArticlesFilteredByContentId)
    api.GET("/user/:uid/articles", controller.GetArticlesByUID)
	api.POST("/article", controller.CreateArticle)

	// User系
	api.POST("/user", controller.CreateUser)

	// history系
	api.GET("/browsedHistory/:user_id", controller.GetBrowsedHistory)
	api.POST("/browsedHistory", controller.CreateBrowsedHistory)

	//api.GET("/oauth", controller.OAuth)

}
