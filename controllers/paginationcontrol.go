package controllers

import (
	"github.com/astaxie/beego"
	"myblog/service"
)

type PaginationController struct {
	beego.Controller
}

var s = services.NewService()
