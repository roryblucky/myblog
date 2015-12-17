package routers

import (
	"github.com/astaxie/beego"
	"myblog/controllers/index"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/category/:categoryid/articles", &controllers.MainController{})
	beego.Router("/page/:pageNum", &controllers.MainController{})

	beego.Router("/article/:id", &controllers.MainController{}, "get:ShowArticle")
}
