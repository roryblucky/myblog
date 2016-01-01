package api

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
	iconFile, fHeader, _ := c.GetFile("icon")
	fullPath := utils.SaveFile(iconFile, fHeader, beego.AppConfig.String("upload_destination"))

	intro := c.Input().Get("introduction")

	blowOwner := models.BlogOwner{Introduction: intro}
	if !utils.IsEmpty(fullPath) {
		blowOwner.ImageIconPath = fullPath
	}
	err := models.UpdateBlogOwner(blowOwner)
	if err != nil {
		errorResult := models.MessageResult{Code: http.StatusInternalServerError, Msg: "Update BlogOwner info failed"}
		c.Data["json"] = &errorResult
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.ServeJson()
		return
	}

	successResult := models.MessageResult{Code: http.StatusOK, Msg: "Update BlogOwner info successful."}
	c.Data["json"] = &successResult
	c.ServeJson()
}
