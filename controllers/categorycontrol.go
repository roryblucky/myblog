package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
	"myblog/service"
	"strconv"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	idStr := c.Ctx.Input.Param(":id")
	pageNum := c.Input().Get("pageNum")

	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			beego.Error(err.Error())
		}
		result, err := services.FindPageRecords(pageNum, &models.Article{}, id)
		if err != nil {
			beego.Error(err)
		}
		if data, ok := result.([]models.Article); ok {
			c.Data["json"] = &data
		}
	} else {
		categories := []models.Category{}
		c.Data["json"] = &categories
		models.DB.Find(&categories)
	}
	c.ServeJson()
}
