package config

import "github.com/zctod/go-tool/config"

// config 配置项
type configure struct {
	Name        string `config:"default:api_demo;comment:项目名称"`
	Host        string `config:"default:127.0.0.1:8080;comment:项目host"`
	SqlDriver   string `config:"default:mysql;comment:数据库驱动（mysql, postgres）"`
	SqlHost     string `config:"default:localhost;comment:数据库地址"`
	SqlPort     string `config:"default:3306;comment:数据库端口"`
	SqlDb       string `config:"default:admin_demo;comment:数据库名称"`
	SqlUsername string `config:"default:root;comment:数据库用户名"`
	SqlPassword string `config:"comment:数据库密码"`
}

var Cfg = &configure{}

// 初始化配置
func Run() {
	// 初始化配置文件
	err := config.InitConfig(Cfg, PATH_ENV)
	if err != nil {
		panic(err)
	}
}
