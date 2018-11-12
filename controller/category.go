package controller

import (
	// "log"
	"net/http"
	"github.com/LINK/service"
	"github.com/LINK/model"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	platformCategories := model.GetPlatformCategories()
	languageCategories := model.GetLanguageCategories()

	rtn := map[string]service.Categories {
		"Platform": platformCategories,
		"Language": languageCategories,
	}
	c.JSON(http.StatusOK, rtn)
	//c.JSON(http.StatusOK, [] map[string]int {
	//    map[string]int{"key1": 10, "key2": 20, "key3": 30},
	//    map[string]int{"key4": 40, "key5": 50, "key6": 60},
	//})
}
