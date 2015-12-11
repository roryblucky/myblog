package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type (
	Article struct {
		Id       int64      // 主键id
		Title    string     `orm:"size(200)"`               // 文章标题
		PostDate time.Time  `orm:"auto_now_add;type(date)"` // 发布时间
		Content  string     `orm:"type(text)"`              // 文章内容
		Category *Category  `orm:"rel(fk)"`                 //分类id
		Comments []*Comment `orm:"reverse(many)"`           //评论
	}

	BlogOwner struct {
		Id            int64  //主键id
		ImageIconPath string `orm:"size(100)"` //照片
		Introduction  string `orm:"size(200)"` //自我介绍
	}

	Category struct {
		Id   int64  //主键id
		Name string `orm:"size(50)"` //分类名称
	}

	Comment struct {
		Id      int64    //主键id
		Content string   `orm:type(text)` //评论内容
		Article *Article `orm:"rel(fk)"`
	}
)

func init() {
	orm.RegisterDataBase("default", "sqlite3", "../data/myblog.db")

	orm.RegisterModel(new(Article), new(BlogOwner), new(Category), new(Comment))
}
