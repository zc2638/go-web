package route

import (
	"api-demo/controller/admin"
	"api-demo/middleware"
	"github.com/gin-gonic/gin"
)

func routeAdmin(g *gin.Engine) {

	admins := g.Group("/admin")
	admins.Use(middleware.AdminAuth)
	admins.GET("/login", new(admin.Auth).Login)
	admins.GET("/logout", new(admin.Auth).Logout)
}