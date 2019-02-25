package controller

import (
	"api-demo/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Base struct {}

func (t *Base) Api(c *gin.Context, sts int, data interface{}, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": sts,
		"data": data,
		"msg": msg,
	})
}

func (t *Base) Succ(c *gin.Context, msg string) {
	t.Api(c, config.CODE_SUCCESS, map[string]interface{}{}, msg)
}

func (t *Base) Data(c *gin.Context, data interface{}) {
	t.Api(c, config.CODE_SUCCESS, data, "")
}

func (t *Base) Err(c *gin.Context, msg string) {
	t.Api(c, config.CODE_FAIL, map[string]interface{}{}, msg)
}