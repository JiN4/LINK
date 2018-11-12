package controller

import (
	"log"
	"net/http"

	"github.com/LINK/model"
	"github.com/LINK/service"
	"github.com/gin-gonic/gin"
)

//POST 閲覧履歴の作成
// type BrowsedHistory struct { //
// 	gorm.Model
// 	UserId    uint
// 	ContentId uint
// }

func CreateBrowsedHistory(c *gin.Context) {
	var browsedHistory service.BrowsedHistory
	c.BindJSON(&browsedHistory)
	log.Println("BIND_JSON: ", browsedHistory)
	model.CreateBrowsedHistory(model.BrowsedHistory{
		UserId:    browsedHistory.UserID,
		ContentId: browsedHistory.ContentID,
	})
	// log.Println(browsedHistory)
	c.JSON(http.StatusOK, browsedHistory)

}

func GetBrowsedHistory(c *gin.Context) {
	userID := c.Param("user_id")
	browsedHistories, err := model.GetEventCardsByBrowsedHistory(userID)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, browsedHistories)
}
