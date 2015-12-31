package routers

import (
	"github.com/astaxie/beego"
	"myblog/controllers/admin"
)

func init() {
	beego.Router("/admin/login", &admin.AdminController{}, "post:Login")
	beego.Router("/admin/logout", &admin.AdminController{}, "get:Logout")

	// angular route
	beego.Router("/admin/main", &admin.AdminController{})
	beego.Router("/admin/article", &admin.AdminController{})
	beego.Router("/admin/article/:op", &admin.AdminController{})
	beego.Router("/admin/category", &admin.AdminController{})
	beego.Router("/admin/blogowner", &admin.AdminController{})
}
