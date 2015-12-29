package api

import (
	"fmt"
	"myblog/models"
	"myblog/utils"
	"net/http"
)

type CategoryController struct {
	AdminBaseController
}

func (c *CategoryController) ShowAllCategories() {
	categories, err := models.GetAllCategories()
	if err != nil {
		errorResult := models.MessageResult{Code: http.StatusInternalServerError, Msg: err.Error()}
		c.Data["json"] = &errorResult
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.ServeJson()
		return
	}
	dataResult := models.DataResult{
		Code: http.StatusOK,
		Data: categories,
	}
	c.Data["json"] = &dataResult
	c.ServeJson()
}

func (c *CategoryController) AddCategory() {
	name := c.Input().Get("name")
	if utils.IsEmpty(name) {
		errorResult := models.MessageResult{Code: http.StatusBadRequest, Msg: "Missing parameter"}
		c.Data["json"] = &errorResult
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.ServeJson()
		return
	}
	err := models.AddCategory(name)
	if err != nil {
		errorResult := models.MessageResult{Code: http.StatusInternalServerError, Msg: err.Error()}
		c.Data["json"] = &errorResult
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.ServeJson()
		return
	}

	successResult := models.MessageResult{Code: http.StatusOK, Msg: fmt.Sprintf("Add categroy succcessful")}
	c.Data["json"] = &successResult
	c.ServeJson()
}

func (c *CategoryController) UpdateCategory() {
	id := c.Ctx.Input.Param(":id")
	if utils.IsEmpty(id) {
		errorResult := models.MessageResult{Code: http.StatusBadRequest, Msg: "Missing parameter"}
		c.Data["json"] = &errorResult
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.ServeJson()
		return
	}

	name := c.Input().Get("name")
	newCategory := models.Category{Name: name}
	err := models.UpdateCategory(id, newCategory)
	if err != nil {
		errorResult := models.MessageResult{Code: http.StatusInternalServerError, Msg: err.Error()}
		c.Data["json"] = &errorResult
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.ServeJson()
		return
	}

	successResult := models.MessageResult{Code: http.StatusOK, Msg: "Update categroy succcessful"}
	c.Data["json"] = &successResult
	c.ServeJson()
}

func (c *CategoryController) DeleteCategory() {
	id := c.Ctx.Input.Param(":id")
	if utils.IsEmpty(id) {
		errorResult := models.MessageResult{Code: http.StatusBadRequest, Msg: "Missing parameter"}
		c.Data["json"] = &errorResult
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.ServeJson()
		return
	}
	err := models.DelCategory(id)
	if err != nil {
		errorResult := models.MessageResult{Code: http.StatusInternalServerError, Msg: err.Error()}
		c.Data["json"] = &errorResult
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.ServeJson()
		return
	}
	successResult := models.MessageResult{Code: http.StatusOK, Msg: fmt.Sprintf("Delete categroy succcessful, id %s", id)}
	c.Data["json"] = &successResult
	c.ServeJson()
}
