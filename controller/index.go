package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Index struct{ Base }

func (ct *Index) Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}
