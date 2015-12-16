package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"myblog/logger"
)

func init() {
	username := beego.AppConfig.String("db_username")
	password := beego.AppConfig.String("db_password")
	host := beego.AppConfig.String("db_host")
	port, err := beego.AppConfig.Int("db_port")
	dbname := beego.AppConfig.String("db_name")

	if err != nil {
		logger.Warn("Port not specify, use default")
		port = 3306
	}
	err = orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, host, port, dbname), 30)
	if err != nil {
		panic(err)
	}
	orm.RegisterModel(new(Comment), new(Category), new(Article), new(BlogOwner))
	orm.Debug = true
}
