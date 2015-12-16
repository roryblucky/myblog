package routers

import (
	"github.com/astaxie/beego"
	"myblog/controllers/index"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Index")
	beego.Router("/article/:id", &controllers.MainController{}, "get:ShowArticle")
}
