package models

type BlogOwner struct {
	ID            int    `json:"id"`    //主键id
	ImageIconPath string `json:"image"` //照片
	Introduction  string `json:"intro"` //自我介绍
}
