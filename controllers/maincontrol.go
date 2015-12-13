package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"myblog/models"
	"myblog/service"
)

type MainController struct {
	beego.Controller
}

var s = services.NewService()

func (c *MainController) Get() {
	//找出所有的信息
	//所属人
	blogOwner := models.BlogOwner{}
	s.Find(&blogOwner)
	//分类
	categories := []models.Category{}
	s.Find(&categories)

	fmt.Println(categories)

	//文章
	result := s.FindPageRecords("1", &models.Article{})
	if a, ok := result.([]models.Article); ok {
		c.Data["articles"] = a
	} else {
		c.Data["articles"] = ""
	}
	c.Data["owner"] = blogOwner
	c.Data["categories"] = categories

	c.TplNames = "index.html"

}
