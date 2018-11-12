package controller

import (
  "fmt"
  "log"
  "net/http"
  "strings"
  "github.com/LINK/model"
  "github.com/gin-gonic/gin"
)

// sort_typeによりSortされたContentsを返す
func GetSortedContentCards(c *gin.Context) {
  fmt.Println("Im On Controller GetSortedContents")
  sort_type := c.Query("sort_type")
  log.Println(sort_type)
  contents, err := model.GetSortedContentCards(sort_type)
  if err != nil {
    log.Println(err)
    c.AbortWithStatus(http.StatusInternalServerError)
  }
  c.JSON(http.StatusOK, contents)
}

func GetContentsFilteredByCategory(c *gin.Context) {
  platformIds := strings.Split(c.Query("platforms"), "-")
  languageIds := strings.Split(c.Query("languages"), "-")
  log.Println("platformIds: \n", platformIds, "\n", languageIds)
  contents, err := model.GetContentsFilteredByCategoryIdList(platformIds, languageIds)
  if err != nil {
    log.Println(err)
    c.AbortWithStatus(http.StatusInternalServerError)
  }
  c.JSON(http.StatusOK, contents)
}

func GetContentById(c *gin.Context) {
  id, err := GetUint(c, "content_id")
  content, err := model.GetContentById(id)
  if err != nil {
    log.Println(err)
    c.AbortWithStatus(http.StatusInternalServerError)
  }
  c.JSON(http.StatusOK, content)
}

func GetAllContentCards(c *gin.Context) {
  contentCards, err := model.GetAllContentCards()
  if err != nil {
    log.Println(err)
    c.AbortWithStatus(http.StatusInternalServerError)
  }
  c.JSON(http.StatusOK, contentCards)
}
