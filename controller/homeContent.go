package controller

import (
	"net/http"

	"github.com/LINK/model"
	"github.com/gin-gonic/gin"
)

func GetHomeContent(c *gin.Context) {
	// uId, err := GetUint(c, "user_id") //URLから記事に結びつくコンテンツのIDを取得
	// if err != nil {
	// 	log.Println(err)
	// 	c.AbortWithStatus(http.StatusInternalServerError)
	// }
	homeContent := model.GetHomeContent()

	c.JSON(http.StatusOK, homeContent)

}
