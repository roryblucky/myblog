package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"myblog/logger"
	"myblog/models"
	"strconv"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	pageNumStr := c.Ctx.Input.Param(":pageNum")
	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil {
		logger.Warn("PageNum not correct, using default")
	}
	articles, hasNextPage, totalPages := models.GetAllArticles(pageNum, 6)

	//如果当前页大于总页数 那么当前页强制等于总页数(最后一页)
	if totalPages < pageNum {
		pageNum = totalPages
	}
	hasPrevPage := false
	if pageNum > 1 {
		hasPrevPage = true
	}

	c.Data["next_page"] = fmt.Sprintf("/page/%d", pageNum+1)
	c.Data["hasNextPage"] = hasNextPage
	c.Data["prev_page"] = fmt.Sprint("/page/%d", pageNum-1)
	c.Data["hasPrevPage"] = hasPrevPage
	c.Data["articles"] = articles

	categories := models.GetAllCategories()
	c.Data["categories"] = categories

	blogOwner := models.GetBlogOwner()
	c.Data["blogOwner"] = blogOwner
	c.TplNames = "index.html"
}
