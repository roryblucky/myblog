package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"myblog/utils"
)

type Comment struct {
	Id      string   `orm:"pk" json:"id"`     //主键id
	Content string   `json:"content"`         //评论内容
	Article *Article `orm:"rel(fk)" json:"-"` //评论主键
}

func AddComment(content, articleId string) (int64, error) {
	o := orm.NewOrm()
	if utils.IsEmpty(content) || utils.IsEmpty(articleId) {
		return 0, errors.New("content or articleid cannot be null")
	}
	comment := Comment{Id: utils.GenerateID(), Content: content, Article: &Article{Id: articleId}}
	return o.Insert(&comment)
}

func DelComment(id string) (int64, error) {
	o := orm.NewOrm()
	if utils.IsEmpty(id) {
		return 0, errors.New("comment id cannot be null")
	}
	comment := Comment{Id: id}
	return o.Delete(&comment)
}
