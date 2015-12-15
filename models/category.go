package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"myblog/utils"
)

type Category struct {
	Id       string     `orm:"pk" json:"id"` //主键id
	Name     string     `json:"name"`        //分类名称
	Articles []*Article `orm:"reverse(many)" json:"articles"`
}

func AddCategory(name string) (int64, error) {
	if utils.IsEmpty(name) {
		return 0, errors.New("category name cannot be null")
	}
	category := &Category{Id: utils.GenerateID(), Name: name}
	o := orm.NewOrm()
	return o.Insert(&category)
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	var categories []*Category
	_, err := o.QueryTable("category").All(&categories)
	return categories, err
}

func DelCategory(id string) (int64, error) {
	category := Category{Id: id}
	o := orm.NewOrm()
	return o.Delete(&category)
}

func UpdateCategory(id string, newCategory Category) error {
	o := orm.NewOrm()
	category := Category{Id: id}
	if o.Read(&category) != nil {
		return errors.New("update category failed")
	}
	category.Name = newCategory.Name
	_, err := o.Update(&category)
	return err
}
