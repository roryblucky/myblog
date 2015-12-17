package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"myblog/logger"
	"myblog/utils"
)

type Category struct {
	Id       string     `orm:"pk" json:"id"` //主键id
	Name     string     `json:"name"`        //分类名称
	Articles []*Article `orm:"reverse(many)" json:"articles"`
}

func AddCategory(name string) error {
	if utils.IsEmpty(name) {
		return errors.New("category name cannot be null")
	}
	utils.DelCache("categories")
	category := Category{Id: utils.GenerateID(), Name: name}
	o := orm.NewOrm()
	_, err := o.Insert(&category)
	return err
}

func GetAllCategories() ([]Category, error) {
	o := orm.NewOrm()
	var categories []Category

	//先从缓存中获取数据
	err := utils.GetDataFromCache("categories", &categories)
	if err != nil {
		_, err = o.QueryTable("category").All(&categories)
		if err != nil {
			logger.Error("Get All Categories failed, error msg: %s", err.Error())
			return nil, err
		}
		utils.SetCache("categories", categories, 900)
	}
	return categories, nil
}

func DelCategory(id string) error {
	utils.DelCache("categories")
	category := Category{Id: id}
	if utils.IsEmpty(id) {
		return errors.New("category id cannot be null")
	}
	o := orm.NewOrm()
	_, err := o.Delete(&category)
	return err
}

func UpdateCategory(id string, newCategory Category) error {
	utils.DelCache("categories")
	o := orm.NewOrm()
	category := Category{Id: id}
	if o.Read(&category) != nil {
		return errors.New("update category failed")
	}
	category.Name = newCategory.Name
	_, err := o.Update(&category)
	return err
}
