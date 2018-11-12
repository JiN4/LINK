package controller

import (
	"log"
	"net/http"

	"github.com/LINK/model"
	"github.com/gin-gonic/gin"
)

func GetAllEventCards(c *gin.Context) {
	events, err := model.GetAllEventCards()
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, events)
}

func GetAllEvents(c *gin.Context) {
  events, err := model.GetAllEvents()
  if err != nil {
    log.Println(err)
    c.AbortWithStatus(http.StatusInternalServerError)
  }
  c.JSON(http.StatusOK, events)
}


func GetEventById(c *gin.Context) {
  id, _ := GetUint(c, "event_id")
  event, err := model.GetEventById(id)
  if err != nil {
    log.Println(err)
    c.AbortWithStatus(http.StatusInternalServerError)
  }
  c.JSON(http.StatusOK, event)
}
