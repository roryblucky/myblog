package admin

import (
	"github.com/astaxie/beego"
	"myblog/models"
	"net/http"
)

type BaseAdminController struct {
	beego.Controller
}

func (c *BaseAdminController) CheckLogin() bool {
	if c.GetSession("admin") == nil {
		errorResult := models.Result{Code: http.StatusUnauthorized, Msg: "Unauthorized"}
		c.Data["json"] = &errorResult
		c.ServeJson()
		return false
	}
	return true
}
