package models

type Comment struct {
	ID        int    `json:"id"`      //主键id
	Content   string `json:"content"` //评论内容
	ArticleId int    `json:"-"`       //评论主键
}
