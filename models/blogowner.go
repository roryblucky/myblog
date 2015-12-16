package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"myblog/logger"
	"myblog/utils"
)

type BlogOwner struct {
	Id            string `orm:"pk" json:"id"` //主键id
	ImageIconPath string `json:"image"`       //照片
	Introduction  string `json:"intro"`       //自我介绍
}

func AddBlogOwner(imagePath, intro string) (int64, error) {
	if utils.IsEmpty(imagePath) || utils.IsEmpty(intro) {
		return 0, errors.New("parameter ImageIcon or Introduction cannot be null")
	}
	blogOwner := BlogOwner{"1", imagePath, intro}
	o := orm.NewOrm()
	return o.Insert(&blogOwner)
}

func GetBlogOwner() (*BlogOwner, error) {
	var owner BlogOwner
	var err error
	err = utils.GetDataFromCache("blogowner", &owner)
	if err != nil {
		logger.Warn(err.Error())
		owner := BlogOwner{Id: "1"}
		o := orm.NewOrm()
		err = o.Read(&owner)
		if err != nil {
			logger.Error("Get Blog owner failed, error msg %s", err.Error())
			return nil, err
		}
		utils.SetCache("blogowner", owner, 9999)
	}
	return &owner, nil
}

func UpdateBlogOwner(b BlogOwner) error {
	utils.DelCache("blogowner")
	o := orm.NewOrm()
	owner := BlogOwner{Id: "1"}
	if o.Read(&owner) == nil {
		owner.ImageIconPath = b.ImageIconPath
		owner.Introduction = b.Introduction
		_, err := o.Update(&owner)
		return err
	}
	return errors.New("update blog owner info failed")
}
