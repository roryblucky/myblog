package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

//http.StatusUnauthorized 401
func (c *ErrorController) Error404() {
	//TODO Add 404 html template
}
