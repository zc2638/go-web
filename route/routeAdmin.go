package route

import (
	"github.com/gin-gonic/gin"
	"go-web/controller/admin"
	"go-web/middleware"
)

func routeAdmin(g *gin.Engine) {

	admins := g.Group("/admin")

	admins.POST("/login", new(admin.Auth).Login)
	admins.GET("/logout", new(admin.Auth).Logout)
	admins.GET("/show", new(admin.Auth).Show)

	admins.Use(middleware.AdminAuth)
	{
		var adminController = new(admin.Admin)
		admins.GET("/admin/list", adminController.List)
		admins.POST("/admin/create", adminController.Create)
		admins.POST("/admin/update", adminController.Update)
		admins.POST("/admin/delete", adminController.Delete)

		admins.GET("/admin/roleList", adminController.RoleList)
		admins.POST("/admin/roleCreate", adminController.RoleCreate)
		admins.POST("/admin/roleUpdate", adminController.RoleUpdate)
		admins.POST("/admin/roleDelete", adminController.RoleDelete)
	}

}