package models

import "time"

type Article struct {
	ID         int       `json:"id"`        // 主键id
	Title      string    `json:"title"`     // 文章标题
	URL        string    `json:"url"`       //文章链接
	PostDate   time.Time `json:"post_date"` // 发布时间
	Content    string    `json:"content"`   // 文章内容
	CategoryId int       `json:"-"`         //分类id
	Comments   []Comment `json:"comments"`  //评论
}
