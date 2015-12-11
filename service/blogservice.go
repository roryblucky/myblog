package services

import (
	"github.com/astaxie/beego/orm"
)

type Service struct {
	orm orm.Ormer
}

func NewService() (s *Service) {
	s = &Service{orm.NewOrm()}
	return
}

func (this *Service) Add(i interface{}) (int64, error) {
	return this.orm.Insert(i)
}

func (this *Service) Remove(i interface{}) (int64, error) {
	return this.orm.Delete(i)
}

func (this *Service) Update(i interface{}, params ...string) (int64, error) {
	return this.orm.Update(i, params...)
}

func (this *Service) Read(i interface{}, params ...string) error {
	return this.orm.Read(i, params...)
}
