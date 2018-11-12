package controller

import (
	// "log"
	"net/http"
	"github.com/LINK/model"

	"github.com/gin-gonic/gin"
)

func GetAllCompanyCards(c *gin.Context) {
	companyCards, _ := model.GetAllCompanyCards()

	c.JSON(http.StatusOK, companyCards)
}
