package admin

import (
	"fmt"
	"myblog/models"
	"myblog/utils"
	"net/http"
)

type CategoryController struct {
	BaseAdminController
}

func (c *CategoryController) ShowAllCategories() {
	isLogin := c.CheckLogin()
	if isLogin {
		categories, err := models.GetAllCategories()
		if err != nil {
			errorResult := models.Result{Code: http.StatusInternalServerError, Msg: err.Error()}
			c.Data["json"] = &errorResult
			c.ServeJson()
			return
		}

		c.Data["json"] = &categories
		c.ServeJson()
	}
}

func (c *CategoryController) AddCategory() {
	isLogin := c.CheckLogin()
	if isLogin {
		name := c.Input().Get("name")
		if utils.IsEmpty(name) {
			errorResult := models.Result{Code: http.StatusBadRequest, Msg: "Missing parameter"}
			c.Data["json"] = &errorResult
			c.ServeJson()
			return
		}
		_, err := models.AddCategory(name)
		if err != nil {
			errorResult := models.Result{Code: http.StatusInternalServerError, Msg: err.Error()}
			c.Data["json"] = &errorResult
			c.ServeJson()
			return
		}

		successResult := models.Result{Code: http.StatusOK, Msg: fmt.Sprintf("Add categroy succcessful")}
		c.Data["json"] = &successResult
		c.ServeJson()
	}
}

func (c *CategoryController) UpdateCategory() {
	isLogin := c.CheckLogin()
	if isLogin {
		id := c.Ctx.Input.Param(":id")
		if utils.IsEmpty(id) {
			errorResult := models.Result{Code: http.StatusBadRequest, Msg: "Missing parameter"}
			c.Data["json"] = &errorResult
			c.ServeJson()
			return
		}

		name := c.Input().Get("name")
		newCategory := models.Category{Name: name}
		err := models.UpdateCategory(id, newCategory)
		if err != nil {
			errorResult := models.Result{Code: http.StatusInternalServerError, Msg: err.Error()}
			c.Data["json"] = &errorResult
			c.ServeJson()
			return
		}

		successResult := models.Result{Code: http.StatusOK, Msg: "Update categroy succcessful"}
		c.Data["json"] = &successResult
		c.ServeJson()
	}
}

func (c *CategoryController) DeleteCategory() {
	isLogin := c.CheckLogin()
	if isLogin {
		id := c.Ctx.Input.Param(":id")
		if utils.IsEmpty(id) {
			errorResult := models.Result{Code: http.StatusBadRequest, Msg: "Missing parameter"}
			c.Data["json"] = &errorResult
			c.ServeJson()
			return
		}
		_, err := models.DelCategory(id)
		if err != nil {
			errorResult := models.Result{Code: http.StatusInternalServerError, Msg: err.Error()}
			c.Data["json"] = &errorResult
			c.ServeJson()
			return
		}
		successResult := models.Result{Code: http.StatusOK, Msg: fmt.Sprintf("Delete categroy succcessful, id %s", id)}
		c.Data["json"] = &successResult
		c.ServeJson()
	}
}