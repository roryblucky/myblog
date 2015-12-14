package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"myblog/models"
	"myblog/service"
	"net/http"
	"strconv"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) GetAllArticles() {
	pageNum := c.Input().Get("pageNum")
	result := services.FindPageRecords(pageNum, &models.Article{})
	if data, ok := result.([]models.Article); ok {
		c.Data["json"] = &data
	}
	c.ServeJson()
}

func (c *ArticleController) GetArticleById() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		errorResult := models.ErrorResult{http.StatusBadRequest, fmt.Sprintf("Bad Request, invalid parameter")}
		c.Data["json"] = &errorResult
	} else {
		article := models.Article{ID: id}
		models.DB.Find(&article).Related(&article.Comments)
		c.Data["json"] = &article
	}
	c.ServeJson()
}

func (c *ArticleController) GetArticlesByCategory() {
	pageNum := c.Input().Get("pageNum")
	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorResult := models.ErrorResult{http.StatusBadRequest, fmt.Sprintf("Bad Request, invalid parameter")}
		c.Data["json"] = &errorResult
	} else {
		result := services.FindPageRecords(pageNum, &models.Article{}, id)
		if data, ok := result.([]models.Article); ok {
			c.Data["json"] = &data
		}
	}
	c.ServeJson()
}
