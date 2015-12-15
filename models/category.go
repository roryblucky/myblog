package models

type Category struct {
	ID       int       `json:"id"`   //主键id
	Name     string    `json:"name"` //分类名称
	Articles []Article `json:"articles"`
}
