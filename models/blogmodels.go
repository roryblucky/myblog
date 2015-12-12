package models

import (
	_ "github.com/mattn/go-sqlite3"
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
