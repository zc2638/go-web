package config

import "github.com/zctod/tool/config"

// config 配置项
type configure struct {
	Name          string `config:"default:lottery"`
	Host          string `config:"default:127.0.0.1:8080"`
	MysqlHost     string `config:"default:localhost"`
	MysqlPort     string `config:"default:3306"`
	MysqlDb       string `config:"default:lottery"`
	MysqlUsername string `config:"default:root"`
	MysqlPassword string
}

var Cfg = &configure{}

// 初始化配置
func Run() {
	// 初始化配置文件
	config.InitConfig(Cfg, PATH_ENV)
}
