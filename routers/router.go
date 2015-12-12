package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"myblog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/page", &controllers.PaginationController{})
}
