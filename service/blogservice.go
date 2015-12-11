package services

import (
	"github.com/astaxie/beego/orm"
)

type Service struct {
	orm.Ormer
}

func NewService() *Service {
	return &Service{orm.NewOrm()}
}

//TODO 分页查询所需方法
