package models

import (
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

type (
	Article struct {
		ID         int       `json:"id"`        // 主键id
		Title      string    `json:"title"`     // 文章标题
		PostDate   time.Time `json:"post_date"` // 发布时间
		Content    string    `json:"content"`   // 文章内容
		CategoryId int       `json:"-"`         //分类id
		Comments   []Comment `json:"comments"`  //评论
	}

	BlogOwner struct {
		ID            int    `json:"id"`    //主键id
		ImageIconPath string `json:"image"` //照片
		Introduction  string `json:"intro"` //自我介绍
	}

	Category struct {
		ID       int       `json:"id"`   //主键id
		Name     string    `json:"name"` //分类名称
		Articles []Article `json:"articles"`
	}

	Comment struct {
		ID        int    `json:"id"`      //主键id
		Content   string `json:"content"` //评论内容
		ArticleId int    `json:"-"`       //评论主键
	}

	ErrorResult struct {
		Code    int    `json:"code"` //状态码
		Message string `json:"msg"`  //消息
	}
)

var DB gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("sqlite3", beego.AppConfig.String("dbpath"))
	if err != nil {
		log.Fatalf("DB initial error ---> %s", err.Error())
		panic(err)
	}
	DB.SingularTable(true)
	DB.LogMode(beego.AppConfig.String("runmode") == "dev")
}
