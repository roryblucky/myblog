package routers

import (
	"github.com/astaxie/beego"
	"myblog/controllers/admin"
)

func init() {
	beego.Router("/admin/login", &admin.AdminController{})
	beego.Router("/admin/blogowner/update", &admin.BlowOwnerController{})
}
