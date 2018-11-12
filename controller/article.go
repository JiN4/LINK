package controller

import (
	"log"
	"net/http"

	"github.com/LINK/model"
	"github.com/gin-gonic/gin"
)

func GetArticlesFilteredByContentId(c *gin.Context) {
	contentId, err := GetUint(c, "content_id") //URLから記事に結びつくコンテンツのIDを取得
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	// article, err := model.GetArticlesFilteredByContentId(contentId)
	articles, err := model.GetArticlesByContentId(contentId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, articles)
}

func GetArticlesByUID(c *gin.Context) {
  uid := c.Param("uid")
  articles, err := model.GetArticlesByUID(uid)
  if err != nil {
    log.Println(err)
    c.AbortWithStatus(http.StatusInternalServerError)
  }
  c.JSON(http.StatusOK, articles)

}

// Articleを追加する
func CreateArticle(c *gin.Context) {
	var article model.Article
	c.BindJSON(&article)

	model.CreateArticle(model.Article{
		ContentID:   article.ContentID,
		Title:       article.Title,
		Description: article.Description,
		Thumbnail:   article.Thumbnail,
		Url:         article.Url,
	})
	log.Println(article)
	c.JSON(http.StatusOK, article)
}
