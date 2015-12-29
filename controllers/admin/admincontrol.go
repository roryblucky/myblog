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
	username := c.Input().Get("username")
	password := c.Input().Get("password")

	if utils.IsEmpty(username) || utils.IsEmpty(password) {
		errorResult := models.MessageResult{Code: http.StatusUnauthorized, Msg: "Empty credentials"}
		c.Data["json"] = &errorResult
		c.ServeJson()
		return
	}
	isSuccess := models.AdminLogin(username, password)
	if !isSuccess {
		errorResult := models.MessageResult{Code: http.StatusUnauthorized, Msg: "Invalid credentials"}
		c.Data["json"] = &errorResult
		c.ServeJson()
		return
	}
	c.SetSession("admin", username)
	result := models.MessageResult{Code: http.StatusOK, Msg: "Login success"}
	c.Data["json"] = &result
	c.ServeJson()
}

// redirect to angular route
func (c *AdminController) Get() {
	if c.GetSession("admin") != nil {
		c.TplNames = "admin/admin_index.html"
	} else {
		c.Redirect("/static/admin/admin_login.html", 302)
	}
}
