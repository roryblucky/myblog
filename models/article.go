package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"myblog/utils"
	"time"
)

type Article struct {
	Id       string     `orm:"pk" json:"id"`                  // 主键id
	Title    string     `json:"title"`                        // 文章标题
	PostDate time.Time  `json:"post_date"`                    // 发布时间
	Content  string     `json:"content"`                      // 文章内容
	Category *Category  `orm:"rel(fk)" json:"-"`              //分类id
	Comments []*Comment `orm:"reverse(many)" json:"comments"` //评论
}

func AddArticle(title, content, categoryId string) (int64, error) {
	o := orm.NewOrm()
	if utils.IsEmpty(title) {
		return 0, errors.New("article title cannot be null")
	}
	newArt := Article{
		Id:       utils.GenerateID(),
		Title:    title,
		PostDate: time.Now().UTC(),
		Content:  content,
		Category: &Category{Id: categoryId},
	}
	return o.Insert(&newArt)
}

func DelArticle(id string) (int64, error) {
	o := orm.NewOrm()
	if utils.IsEmpty(id) {
		return 0, errors.New("article id cannot be null")
	}
	art := Article{Id: id}
	return o.Delete(&art)
}

func UpdateArticle(id string, newArt Article) error {
	o := orm.NewOrm()
	art := Article{Id: id}
	if o.Read(&art) != nil {
		return errors.New("update article failed")
	}
	art.Category = newArt.Category
	art.Content = newArt.Content
	art.Title = newArt.Title
	_, err := o.Update(&art)
	return err
}

func GetArticleById(id string) (Article, error) {
	o := orm.NewOrm()
	art := Article{Id: id}
	err := o.Read(&art)
	o.QueryTable("comment").Filter("article_id", id).RelatedSel().All(&art.Comments)
	return art, err
}
