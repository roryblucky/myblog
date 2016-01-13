package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"myblog/logger"
	"myblog/utils"
	"time"
)

type Article struct {
	Id       string     `orm:"pk" json:"id"`                  // 主键id
	Title    string     `json:"title"`                        // 文章标题
	PostDate time.Time  `json:"post_date"`                    // 发布时间
	Content  string     `json:"content"`                      // 文章内容
	Category *Category  `orm:"rel(fk)" json:"category"`       // 分类id
	Comments []*Comment `orm:"reverse(many)" json:"comments"` // 评论
}

func AddArticle(title, content, categoryId string) error {
	o := orm.NewOrm()
	if utils.IsEmpty(title) {
		return errors.New("article title cannot be null")
	}
	newArt := Article{
		Id:       utils.GenerateID(),
		Title:    title,
		PostDate: time.Now().UTC(),
		Content:  content,
		Category: &Category{Id: categoryId},
	}
	_, err := o.Insert(&newArt)
	return err
}

func DelArticle(id string) error {
	o := orm.NewOrm()
	if utils.IsEmpty(id) {
		return errors.New("article id cannot be null")
	}
	utils.DelCache(fmt.Sprintf("blog-article-%s", id))
	art := Article{Id: id}
	_, err := o.Delete(&art)
	return err
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
	utils.DelCache(fmt.Sprintf("blog-article-%s", id))
	_, err := o.Update(&art)
	return err
}

func GetArticleById(id string) (Article, error) {
	art := Article{Id: id}
	err := utils.GetDataFromCache(fmt.Sprintf("blog-article-%s", id), &art)
	if err != nil {
		o := orm.NewOrm()
		err = o.Read(&art)
		if err == nil {
			o.QueryTable("comment").Filter("article_id", id).RelatedSel().All(&art.Comments)
			utils.SetCache(fmt.Sprintf("blog-article-%s", id), art, 300)
		}
	}
	return art, err
}

func GetAllArticles(pageNum, pageSize int) (bool, int, []Article) {
	return GetArticlesByCondition(pageNum, pageSize, nil)
}

//pageNum 当前页 pageSize 每页条数
func GetArticlesByCondition(pageNum, pageSize int, condition interface{}) (bool, int, []Article) {
	if pageNum < 1 {
		pageNum = 1
	}

	if pageSize < 1 {
		pageSize = 6
	}
	var articles []Article
	// 根据分类获取所有的文章
	o := orm.NewOrm()

	qs := o.QueryTable("article").Limit(pageSize).Offset(pageSize * (pageNum - 1))

	switch condition.(type) {
	case *Category:
		category := condition.(*Category)
		qs = qs.Filter("category_id", category.Id).RelatedSel()
		qs.All(&articles, "Id", "Title", "PostDate", "Content")

	default:
		qs.All(&articles)
	}
	// 计算总共有多少条记录
	num, err := qs.Count()
	totalRecords := int(num)
	if err != nil {
		logger.Warn(err.Error())
		totalRecords = 0
	}
	// 计算总共有多少页
	totalPages := totalRecords / pageSize
	if totalPages == 0 || totalPages%pageSize == 0 {
		totalPages = 1
	} else if totalPages != 0 {
		totalPages += 1
	}

	// 是否有下一页
	hasNextPage := true
	if pageNum == totalPages {
		hasNextPage = false
	}
	return hasNextPage, totalPages, articles
}
