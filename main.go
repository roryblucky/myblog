package main

import (
	"github.com/astaxie/beego"
	_ "myblog/routers"
	"myblog/utils"
)

func main() {
	beego.AddFuncMap("blogDateFormat", utils.DateFormat)
	beego.Run()
}
