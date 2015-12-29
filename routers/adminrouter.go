package routers

import (
	"github.com/astaxie/beego"
	"myblog/controllers/admin"
)

func init() {
	beego.Router("/admin/login", &admin.AdminController{})
	beego.Router("/api/admin/blogowner/update", &admin.BlowOwnerController{})

	beego.Router("/api/admin/categories", &admin.CategoryController{}, "get:ShowAllCategories")
	beego.Router("/api/admin/category/add", &admin.CategoryController{}, "post:AddCategory")
	beego.Router("/api/admin/category/update/:id", &admin.CategoryController{}, "post:UpdateCategory")
	beego.Router("/api/admin/category/del/:id", &admin.CategoryController{}, "get:DeleteCategory")

	beego.Router("/api/admin/articles", &admin.ArticleController{}, "get:GetArticles")
	beego.Router("/api/admin/article/del/:id", &admin.ArticleController{}, "post:DelArticle")
	beego.Router("/api/admin/article/update/:id", &admin.ArticleController{}, "post:UpdateArticle")
	beego.Router("/api/admin/article/add", &admin.ArticleController{}, "post:AddArticle")
}
