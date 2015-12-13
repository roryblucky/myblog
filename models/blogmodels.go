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
		ID       int        // 主键id
		Title    string     // 文章标题
		PostDate time.Time  // 发布时间
		Content  string     // 文章内容
		Category *Category  //分类id
		Comments []*Comment //评论
	}

	BlogOwner struct {
		ID            int    //主键id
		ImageIconPath string //照片
		Introduction  string //自我介绍
	}

	Category struct {
		ID       int    //主键id
		Name     string //分类名称
		Articles []*Article
	}

	Comment struct {
		ID      int    //主键id
		Content string //评论内容
	}
)

var DB gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("sqlite3", beego.AppConfig.String("dbpath"))
	if err != nil {
		log.Fatalf("DB initial error ---> %s", err.Error())
	}
	DB.SingularTable(true)
	DB.LogMode(beego.AppConfig.String("runmode") == "dev")
}
