package admin

import (
	"github.com/astaxie/beego"
	"myblog/models"
	"myblog/utils"
	"net/http"
)

type BlowOwnerController struct {
	AdminBaseController
}

func (c *BlowOwnerController) Post() {
	isLogin := c.CheckLogin()
	if isLogin {
		iconFile, fHeader, _ := c.GetFile("icon")
		fullPath := utils.SaveFile(iconFile, fHeader, beego.AppConfig.String("upload_destination"))

		if utils.IsEmpty(fullPath) {
			errorResult := models.MessageResult{Code: http.StatusInternalServerError, Msg: "Upload failed"}
			c.Data["json"] = &errorResult
			c.ServeJson()
			return
		}

		intro := c.Input().Get("introduction")
		blowOwner := models.BlogOwner{ImageIconPath: fullPath, Introduction: intro}
		err := models.UpdateBlogOwner(blowOwner)
		if err != nil {
			errorResult := models.MessageResult{Code: http.StatusInternalServerError, Msg: "Update BlogOwner info failed"}
			c.Data["json"] = &errorResult
			c.ServeJson()
			return
		}

		successResult := models.MessageResult{Code: http.StatusOK, Msg: "Update BlogOwner info successful."}
		c.Data["json"] = &successResult
		c.ServeJson()
	}
}
