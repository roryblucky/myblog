package admin

import (
	"github.com/astaxie/beego"
	"myblog/models"
	"myblog/utils"
	"net/http"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Post() {
	userName := c.Input().Get("userName")
	password := c.Input().Get("password")

	if utils.IsEmpty(userName) || utils.IsEmpty(password) {
		errorResult := models.MessageResult{Code: http.StatusForbidden, Msg: "Empty credentials"}
		c.Data["json"] = &errorResult
		c.ServeJson()
		return
	}

	isSuccess := models.AdminLogin(userName, password)
	if !isSuccess {
		errorResult := models.MessageResult{Code: http.StatusForbidden, Msg: "Invalid credentials"}
		c.Data["json"] = &errorResult
		c.ServeJson()
		return
	}
	c.SetSession("admin", userName)
	result := models.MessageResult{Code: http.StatusOK, Msg: "Login success"}
	c.Data["json"] = &result
	c.ServeJson()
}
