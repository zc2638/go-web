package middleware

import (
	"api-demo/config"
	"api-demo/lib/database"
	"api-demo/lib/jwt"
	"api-demo/model"
	"github.com/gin-gonic/gin"
	"github.com/zctod/tool/common/utils"
	"net/http"
	"strings"
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
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": config.CODE_FAIL,
			"msg": "登陆失败",
		})
		return
	}
	db, err := database.Open()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": config.CODE_FAIL,
			"msg": "登陆失败，系统错误",
		})
		return
	}
	jwtData, err := jwt.ParseInfo(tokenStr, config.JWT_SECRET_ADMIN)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": config.CODE_FAIL,
			"msg": "异常登录信息",
		})
		return
	}
	info, ok := jwtData["info"]
	if !ok {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": config.CODE_FAIL,
			"msg": "异常登录信息",
		})
		return
	}
	roleId, ok := info.(map[string]interface{})["role"]
	if !ok {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": config.CODE_FAIL,
			"msg": "异常登录信息&",
		})
		return
	}

	role := model.AdminRole{}
	db.Where("id", roleId).First(&role)
	rule := strings.Split(role.Rule, ",")
	if exist, _ := utils.InArray(c.Request.URL, rule); !exist {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": config.CODE_FAIL,
			"msg": "无此权限",
		})
		return
	}
	c.Set("admin", jwtData)

	c.Next()
}