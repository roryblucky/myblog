package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@/my_blog?charset=utf8", 30)
	orm.RegisterModel(new(Comment), new(Category), new(Article), new(BlogOwner))
	orm.Debug = true
}
