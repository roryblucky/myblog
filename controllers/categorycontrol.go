package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) GetAllCategories() {
	categories := []models.Category{}
	models.DB.Find(&categories)
	c.Data["json"] = &categories
	c.ServeJson()
}
