package main

import (
	"github.com/astaxie/beego"
	_ "myblog/routers"
)

func main() {
	beego.AutoRender = false // 关闭自动模版渲染
	beego.Run()
}
