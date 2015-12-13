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
	//找出所有的信息
	//所属人
	blogOwner := models.BlogOwner{}
	models.DB.Find(&blogOwner)
	//分类
	categories := []models.Category{}
	models.DB.Find(&categories)

	//文章
	result, err := services.FindPageRecords("1", &models.Article{})
	utils.LogError(err)

	if data, ok := result.([]models.Article); ok {
		c.Data["articles"] = data
	} else {
		c.Data["articles"] = ""
	}
	c.Data["owner"] = blogOwner
	c.Data["categories"] = categories
	c.Data["isIndex"] = true
	c.Data["title"] = "主页"
	c.TplNames = "index.html"
}
