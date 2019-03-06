package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zctod/tool/common/util_server"
	"go-web/controller"
	"go-web/middleware"
	"net/http"
	"time"
)

func Run() {

	var g = gin.Default()
	g.Use(middleware.Cors())

	// 加载静态文件
	g.Static("/public", "./public")
	g.GET("/", new(controller.Index).Index)

	// 注册路由
	routeApi(g)
	routeAdmin(g)
	//开启服务
	startServer(g)
}

func startServer(g *gin.Engine) {

	server := &http.Server{
		Addr:           ":8080",
		Handler:        g,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()

	// 平滑退出，先结束所有在执行的任务
	util_server.GracefulExitWeb(server)
}
