package routers

import (
	"github.com/astaxie/beego"
	"myblog/controllers/admin/api"
)

func init() {
	beego.Router("/api/admin/blogowner/update", &api.BlowOwnerController{})

	beego.Router("/api/admin/categories", &api.CategoryController{}, "get:ShowAllCategories")
	beego.Router("/api/admin/category/add", &api.CategoryController{}, "post:AddCategory")
	beego.Router("/api/admin/category/update/:id", &api.CategoryController{}, "post:UpdateCategory")
	beego.Router("/api/admin/category/del/:id", &api.CategoryController{}, "get:DeleteCategory")

	beego.Router("/api/admin/articles", &api.ArticleController{}, "get:GetArticles")
	beego.Router("/api/admin/article/del/:id", &api.ArticleController{}, "post:DelArticle")
	beego.Router("/api/admin/article/update/:id", &api.ArticleController{}, "post:UpdateArticle")
	beego.Router("/api/admin/article/add", &api.ArticleController{}, "post:AddArticle")
}
