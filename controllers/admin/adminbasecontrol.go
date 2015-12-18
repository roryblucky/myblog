package admin

import (
	"github.com/astaxie/beego"
	"myblog/models"
	"net/http"
)

type AdminBaseController struct {
	beego.Controller
}

func (c *AdminBaseController) Prepare() {
	if c.GetSession("admin") == nil {
		errorResult := models.MessageResult{Code: http.StatusUnauthorized, Msg: "Unauthorized"}
		c.Data["json"] = &errorResult
		c.ServeJson()
		return
	}
}
