package models

import "github.com/astaxie/beego/orm"

type AdminUser struct {
	Id       string `orm:"pk"`
	UserName string
	Password string
}

func AdminLogin(userName, password string) bool {
	user := AdminUser{UserName: userName, Password: password}
	o := orm.NewOrm()
	if o.Read(&user) == nil {
		return true
	}
	return false
}
