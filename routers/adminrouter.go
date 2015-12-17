package routers

import (
	"github.com/astaxie/beego"
	"myblog/controllers/admin"
)

func init() {
	beego.Router("/admin/login", &admin.AdminController{})
	beego.Router("/admin/blogowner/update", &admin.BlowOwnerController{})
	beego.Router("/admin/categories", &admin.CategoryController{}, "get:ShowAllCategories")
	beego.Router("/admin/category/add", &admin.CategoryController{}, "post:AddCategory")
	beego.Router("/admin/category/update/:id", &admin.CategoryController{}, "post:UpdateCategory")
	beego.Router("/admin/category/del/:id", &admin.CategoryController{}, "get:DeleteCategory")

}
