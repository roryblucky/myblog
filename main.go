package main

import (
	"github.com/astaxie/beego"
	"myblog/controllers"
	_ "myblog/routers"
	"myblog/utils"
)

func main() {
	beego.AddFuncMap("blogDateFormat", utils.DateFormat)
	beego.ErrorController(&controllers.ErrorController{})
	//tes tag
	beego.Run()
}
