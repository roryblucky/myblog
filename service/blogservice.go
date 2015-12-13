package services

import (
	"container/list"
	"github.com/astaxie/beego"
	_ "github.com/mattn/go-sqlite3"
	"myblog/models"
	"strconv"
)

type Pager struct {
	currentPage  int         //当前页
	totalPages   int         //总共多少页
	totalRecords int         //总记录数
	startIndex   int         //开始位置索引
	pageSize     int         //每页多少条数据
	Records      interface{} //数据
}

func newPager(currentPage, totalRecords int) (*Pager, error) {
	pageSize, err := strconv.Atoi(beego.AppConfig.String("pagesize"))
	startIndex := (currentPage - 1) * pageSize
	totalPages := 0
	if totalRecords%pageSize == 0 {
		totalPages = totalRecords / pageSize
	} else {
		totalPages = totalRecords/pageSize + 1
	}
	return &Pager{
		currentPage:  1,
		totalPages:   totalPages,
		totalRecords: totalRecords,
		startIndex:   startIndex,
		pageSize:     pageSize,
		Records:      list.New(),
	}, err
}

func GetTotalRecords(i interface{}) (num int) {
	models.DB.Model(i).Count(&num)
	return
}

func FindPageRecords(pageNum string, i interface{}) (result interface{}, err error) {
	currentPage := 0
	if pageNum != "" {
		currentPage, err = strconv.Atoi(pageNum)
	}
	var pager *Pager
	pager, err = newPager(currentPage, GetTotalRecords(i))

	switch i.(type) {
	case *models.BlogOwner:
		data := []models.BlogOwner{}
		models.DB.Limit(pager.pageSize).Offset(pager.startIndex).Find(&data)
		result = data
	case *models.Article:
		data := []models.Article{}
		models.DB.Limit(pager.pageSize).Offset(pager.startIndex).Find(&data)
		result = data
	case *models.Category:
		data := []models.Category{}
		models.DB.Limit(pager.pageSize).Offset(pager.startIndex).Find(&data)
		result = data
	case *models.Comment:
		data := []models.Comment{}
		models.DB.Limit(pager.pageSize).Offset(pager.startIndex).Find(&data)
		result = data
	}
	return
}
