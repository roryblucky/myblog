package api

import (
	"fmt"
	"myblog/logger"
	"myblog/models"
	"myblog/utils"
	"net/http"
	"strconv"
)

type ArticleController struct {
	AdminBaseController
}

func (c *ArticleController) AddArticle() {
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	categoryId := c.Input().Get("category_id")
	if utils.IsEmpty(categoryId) || utils.IsEmpty(title) {
		errorResult := models.MessageResult{Code: http.StatusBadRequest, Msg: "Title or Category cannot be empty"}
		c.Data["json"] = &errorResult
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.ServeJson()
		return
	}

	err := models.AddArticle(title, content, categoryId)
	if err != nil {
		errorResult := models.MessageResult{
			Code: http.StatusInternalServerError,
			Msg:  fmt.Sprintf("Add article failed, error msg: %s", err.Error()),
		}
		c.Data["json"] = &errorResult
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.ServeJson()
		return
	}
	successResult := models.MessageResult{Code: http.StatusOK, Msg: "Add Article successful"}
	c.Data["json"] = &successResult
	c.ServeJson()
}

func (c *ArticleController) UpdateArticle() {
	id := c.Ctx.Input.Param(":id")
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	categoryId := c.Input().Get("category_id")
	if utils.IsEmpty(id) || utils.IsEmpty(categoryId) || utils.IsEmpty(title) {
		errorResult := models.MessageResult{Code: http.StatusBadRequest, Msg: "Title or Category cannot be empty"}
		c.Data["json"] = &errorResult
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.ServeJson()
		return
	}
	newArticle := models.Article{Id: id, Title: title, Content: content, Category: &models.Category{Id: categoryId}}

	err := models.UpdateArticle(id, newArticle)
	if err != nil {
		errorResult := models.MessageResult{
			Code: http.StatusInternalServerError,
			Msg:  fmt.Sprintf("Update article failed, error msg: %s", err.Error()),
		}
		c.Data["json"] = &errorResult
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.ServeJson()
		return
	}
	successResult := models.MessageResult{Code: http.StatusOK, Msg: "Update Article successful"}
	c.Data["json"] = &successResult
	c.ServeJson()
}

func (c *ArticleController) DelArticle() {
	id := c.Ctx.Input.Param(":id")
	if utils.IsEmpty(id) {
		errorResult := models.MessageResult{Code: http.StatusBadRequest, Msg: "Id cannot be empty"}
		c.Data["json"] = &errorResult
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.ServeJson()
		return
	}
	err := models.DelArticle(id)
	if err != nil {
		errorResult := models.MessageResult{
			Code: http.StatusInternalServerError,
			Msg:  fmt.Sprintf("Delete article failed, error msg: %s", err.Error()),
		}
		c.Data["json"] = &errorResult
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.ServeJson()
		return
	}
	successResult := models.MessageResult{Code: http.StatusOK, Msg: "Update Article successful"}
	c.Data["json"] = &successResult
	c.ServeJson()
}

func (c *ArticleController) GetArticleById() {
	id := c.Ctx.Input.Param(":id")
	if utils.IsEmpty(id) {
		errorResult := models.MessageResult{Code: http.StatusBadRequest, Msg: "Id cannot be empty"}
		c.Data["json"] = &errorResult
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.ServeJson()
		return
	}

	article, err := models.GetArticleById(id)
	if err != nil {
		errorResult := models.MessageResult{Code: http.StatusNotFound, Msg: "Result cannot be found"}
		c.Data["json"] = &errorResult
		c.Ctx.Output.SetStatus(http.StatusNotFound)
		c.ServeJson()
		return
	}
	successResult := models.DataResult{
		Code: http.StatusOK,
		Data: article,
	}
	c.Data["json"] = &successResult
	c.ServeJson()

}

func (c *ArticleController) GetArticles() {
	pageNumStr := c.Ctx.Input.Param(":pageNum")
	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil {
		logger.Warn("PageNum not correct, using default")
		pageNum = 0
	}
	hasNextPage, totalPages, articles := models.GetAllArticles(pageNum, 6)

	//如果当前页大于总页数 那么当前页强制等于总页数(最后一页)
	if totalPages < pageNum {
		pageNum = totalPages
	}
	hasPrevPage := false
	if pageNum > 1 {
		hasPrevPage = true
	}
	dataResult := models.DataResult{
		Code:        http.StatusOK,
		NextPage:    fmt.Sprintf("/page/%d", pageNum+1),
		HasNextPage: hasNextPage,
		PrevPage:    fmt.Sprintf("/page/%d", pageNum-1),
		HasPrePage:  hasPrevPage,
		Data:        articles,
	}
	c.Data["json"] = &dataResult
	c.ServeJson()
}
