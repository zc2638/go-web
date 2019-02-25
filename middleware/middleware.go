package middleware

import (
	"api-demo/config"
	"api-demo/lib/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 跨域支持
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func AdminAuth(c *gin.Context) {

	tokenStr := c.Request.Header.Get("token")
	if jwt.CheckValid(tokenStr, config.JWT_SECRET_ADMIN) == false {
		c.JSON(http.StatusOK, gin.H{
			"code": config.CODE_FAIL,
			"msg": "登陆失败",
		})
		return
	}
	c.Next()
}