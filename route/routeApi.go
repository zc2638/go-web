package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func routeApi(g *gin.Engine) {

	apis := g.Group("/api")
	apis.GET("/index", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
}