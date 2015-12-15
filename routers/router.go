package routers

import (
	"github.com/astaxie/beego"
	"myblog/controllers/index"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
