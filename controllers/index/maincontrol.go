package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"myblog/logger"
	"myblog/models"
	"myblog/utils"
	"net/http"
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
		pageNum = 1
	}
	categoryId := c.Ctx.Input.Param(":categoryid")

	var articles []models.Article
	hasNextPage := false
	totalPages := 0
	if utils.IsEmpty(categoryId) {
		hasNextPage, totalPages, articles = models.GetAllArticles(pageNum, 6)
	} else {
		category := models.Category{Id: categoryId}
		hasNextPage, totalPages, articles = models.GetArticlesByCondition(pageNum, 6, &category)
	}

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
	c.Data["prev_page"] = fmt.Sprintf("/page/%d", pageNum-1)
	c.Data["hasPrevPage"] = hasPrevPage
	c.Data["articles"] = articles
	categories, _ := models.GetAllCategories()
	c.Data["categories"] = categories

	blogOwner, _ := models.GetBlogOwner()
	c.Data["blogOwner"] = blogOwner
	c.Data["title"] = "首页"
	c.TplNames = "index.html"
}

func (c *MainController) ShowArticle() {
	id := c.Ctx.Input.Param(":id")

	article, err := models.GetArticleById(id)
	if err != nil {
		c.Abort(strconv.Itoa(http.StatusNotFound))
	}
	c.Data["article"] = &article

	categories, _ := models.GetAllCategories()
	c.Data["categories"] = categories

	blogOwner, _ := models.GetBlogOwner()
	c.Data["blogOwner"] = blogOwner

	c.TplNames = "article.html"
}
