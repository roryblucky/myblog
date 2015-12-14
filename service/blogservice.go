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

func newPager(currentPage, totalRecords int) *Pager {
	pageSize, _ := strconv.Atoi(beego.AppConfig.String("pagesize"))
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
	}
}

func GetTotalRecords(i interface{}) (num int) {
	models.DB.Model(i).Count(&num)
	return
}

func FindPageRecords(pageNum string, i interface{}, where ...interface{}) (result interface{}) {
	currentPage := 1
	if pageNum != "" {
		num, _ := strconv.Atoi(pageNum)
		if num != 0 {
			currentPage = num
		}
	}
	var pager *Pager
	pager = newPager(currentPage, GetTotalRecords(i))

	switch i.(type) {
	case *models.BlogOwner:
		data := []models.BlogOwner{}
		models.DB.Limit(pager.pageSize).Offset(pager.startIndex).Find(&data)
		result = data
	case *models.Article:
		articles := []models.Article{}
		if where != nil {
			models.DB.Where("category_id = ?", where).Limit(pager.pageSize).Offset(pager.startIndex).Find(&articles)
		} else {
			models.DB.Limit(pager.pageSize).Offset(pager.startIndex).Find(&articles)
		}

		for i := 0; i < len(articles); i++ {
			models.DB.Model(&articles[i]).Related(&articles[i].Comments)
		}
		result = articles
	case *models.Category:
		categories := []models.Category{}
		models.DB.Limit(pager.pageSize).Offset(pager.startIndex).Find(&categories)

		for i := 0; i < len(categories); i++ {
			models.DB.Model(&categories[i]).Related(&categories[i].Articles)
		}
		result = categories
	case *models.Comment:
		data := []models.Comment{}
		models.DB.Limit(pager.pageSize).Offset(pager.startIndex).Find(&data)
		result = data
	}
	return
}
