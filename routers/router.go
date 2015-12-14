package routers

import (
	"github.com/astaxie/beego"
	"myblog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/category/?:id", &controllers.CategoryController{})
	beego.Router("/page", &controllers.PaginationController{})
}
