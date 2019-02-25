package database

import (
	"api-demo/config"
	"api-demo/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Open() (*gorm.DB, error) {
	cfg := config.Cfg
	return gorm.Open("mysql", cfg.MysqlUsername+":"+cfg.MysqlPassword+"@tcp("+cfg.MysqlHost+")/"+cfg.MysqlDb+"?charset=utf8mb4&parseTime=True&loc=Local")
}

func DBMigrate() {

	db, err := Open()
	if err != nil {
		panic("open sql error:" + err.Error())
	}

	// 禁用表名复数
	//db.SingularTable(true)

	// 自动生成表结构
	db.AutoMigrate(
		&model.Admin{},
	)
}