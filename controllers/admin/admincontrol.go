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

func (c *AdminController) Login() {
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

func (c *AdminController) Logout() {
	c.DelSession("admin")
	c.Redirect("/static/admin/admin_login.html", 302)
}

// redirect to angular route
func (c *AdminController) Get() {
	if name, ok := c.GetSession("admin").(string); ok {
		c.TplNames = "admin/admin_index.html"
		c.Data["name"] = name
	} else {
		c.Redirect("/static/admin/admin_login.html", 302)
	}
}
