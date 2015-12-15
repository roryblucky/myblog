package models

import (
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
