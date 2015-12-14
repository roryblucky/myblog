package routers

import (
	"github.com/astaxie/beego"
	"myblog/controllers"
)

func init() {
	beego.Router("/articles", &controllers.ArticleController{}, "get:GetAllArticles")
	beego.Router("/articles/category/:id", &controllers.ArticleController{}, "get:GetArticlesByCategory")
	beego.Router("/article/:id", &controllers.ArticleController{}, "get:GetArticleById")

	beego.Router("/categories", &controllers.CategoryController{}, "get:GetAllCategories")
}
