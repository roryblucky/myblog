package services

import (
	"container/list"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"myblog/models"
	"myblog/utils"
	"strconv"
)

type Service struct {
	gorm.DB
}

func NewService() *Service {
	db, err := gorm.Open("sqlite3", beego.AppConfig.String("dbpath"))
	utils.HandleErr(err)
	db.SingularTable(true)
	db.LogMode(true)
	return &Service{db}
}

type Pager struct {
	currentPage  int         //当前页
	totalPages   int         //总共多少页
	totalRecords int         //总记录数
	startIndex   int         //开始位置索引
	pageSize     int         //每页多少条数据
	Records      interface{} //数据
}

func newPager(currentPage, totalRecords int) *Pager {
	pageSize := 6
	startIndex := (currentPage - 1) * pageSize
	totalPages := 0
	if totalRecords%pageSize == 0 {
		totalPages = totalRecords / pageSize
	} else {
		totalPages = totalRecords/pageSize + 1
	}
	return &Pager{currentPage: 1, totalPages: totalPages, totalRecords: totalRecords, startIndex: startIndex, pageSize: pageSize, Records: list.New()}
}

func (this *Service) GetTotalRecords(i interface{}) (num int) {
	this.Model(i).Count(&num)
	return
}

func (this *Service) FindPageRecords(pageNum string, i interface{}) (result interface{}) {
	currentPage := 0
	if pageNum != "" {
		pageNum, _ := strconv.Atoi(pageNum)
		currentPage = pageNum
	}
	pager := newPager(currentPage, this.GetTotalRecords(i))

	switch i.(type) {
	case *models.BlogOwner:
		data := []models.BlogOwner{}
		this.Limit(pager.pageSize).Offset(pager.startIndex).Find(&data)
		result = data
	case *models.Article:
		data := []models.Article{}
		this.Limit(pager.pageSize).Offset(pager.startIndex).Find(&data)
		result = data
	case *models.Category:
		data := []models.Category{}
		this.Limit(pager.pageSize).Offset(pager.startIndex).Find(&data)
		result = data
	case *models.Comment:
		data := []models.Comment{}
		this.Limit(pager.pageSize).Offset(pager.startIndex).Find(&data)
		result = data
	}
	return pager
}
