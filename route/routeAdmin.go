package route

import (
	"api-demo/controller/admin"
	"api-demo/middleware"
	"github.com/gin-gonic/gin"
)

func routeAdmin(g *gin.Engine) {

	admins := g.Group("/admin")

	admins.GET("/login", new(admin.Auth).Login)
	admins.GET("/logout", new(admin.Auth).Logout)

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