package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-web/config"
	"go-web/model"
)

func Open() (*gorm.DB, error) {
	cfg := config.Cfg
	var optionStr string
	switch cfg.SqlDriver {
	case "mysql":
		optionStr = cfg.SqlUsername + ":" + cfg.SqlPassword + "@tcp(" + cfg.SqlHost + ":" + cfg.SqlPort + ")/" + cfg.SqlDb + "?charset=utf8mb4&parseTime=True&loc=Local"
	case "postgres":
		optionStr = "host=" + cfg.SqlHost + ":" + cfg.SqlPort + " user=" + cfg.SqlUsername + " dbname=" + cfg.SqlDb + " sslmode=disable password=" + cfg.SqlPassword
		break
	default:
		break
	}
	return gorm.Open(cfg.SqlDriver, optionStr)
}

func DBMigrate() {

	db, err := Open()
	defer db.Close()
	if err != nil {
		fmt.Println("open sql error:" + err.Error())
	}

	// 禁用表名复数
	//db.SingularTable(true)

	// 自动生成表结构
	db.AutoMigrate(
		&model.Admin{},
		&model.AdminRole{},
	)
}