package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
	"myblog/service"
	"myblog/utils"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	result, err := services.FindPageRecords("1", &models.Article{})
	utils.LogError(err)

	if data, ok := result.([]models.Article); ok {
		c.Data["json"] = &data
	}
	c.ServeJson()
}
