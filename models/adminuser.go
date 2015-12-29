package models

import (
	"github.com/astaxie/beego/orm"
)

type AdminUser struct {
	Id       string `orm:"pk"`
	Username string
	Password string
}

func AdminLogin(username, password string) bool {
	o := orm.NewOrm()
	isExist := o.QueryTable("admin_user").Filter("username", username).Filter("password", password).Exist()
	if isExist {
		return true
	}
	return false
}
