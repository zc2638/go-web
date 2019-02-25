package main

import (
	"api-demo/config"
	"api-demo/lib/database"
	"api-demo/route"
)

// 数据库migrate: gorm
// 数据库操作: gorose
// http框架: gin
func main() {

	// 初始化配置
	config.Run()

	// 生成数据库表结构
	database.DBMigrate()

	// 注册http服务
	route.Run()
}