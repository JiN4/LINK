package router

import (
	"net/http"
	_ "net/http"

	"github.com/gin-gonic/gin"
)

// GetRouter ルーターを設定したgin.Engineを返す

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/alive", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	api := r.Group("/api/v1")
	apiRouter(api)

	return r
}
