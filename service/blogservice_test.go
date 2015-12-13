package services

import (
	"myblog/models"
	"testing"
)

func TestGetTotalRecords(t *testing.T) {
	blogOwner1 := &models.BlogOwner{Introduction: "this is the test1", ImageIconPath: "this is the test"}
	blogOwner2 := &models.BlogOwner{Introduction: "this is the test2", ImageIconPath: "this is the test"}
	blogOwner3 := &models.BlogOwner{Introduction: "this is the test3", ImageIconPath: "this is the test"}
	blogOwner4 := &models.BlogOwner{Introduction: "this is the test4", ImageIconPath: "this is the test"}

	models.DB.Create(blogOwner1)
	models.DB.Create(blogOwner2)
	models.DB.Create(blogOwner3)
	models.DB.Create(blogOwner4)

	num := GetTotalRecords(models.BlogOwner{})
	if num != 4 {
		t.Fail()
		t.Fatal("Actual size is ", num)
	}

	models.DB.Delete(blogOwner1)
	models.DB.Delete(blogOwner2)
	models.DB.Delete(blogOwner3)
	models.DB.Delete(blogOwner4)

}

func TestFindPageRecords(t *testing.T) {
	blogOwner1 := &models.BlogOwner{Introduction: "this is the test1", ImageIconPath: "this is the test"}
	blogOwner2 := &models.BlogOwner{Introduction: "this is the test2", ImageIconPath: "this is the test"}
	blogOwner3 := &models.BlogOwner{Introduction: "this is the test3", ImageIconPath: "this is the test"}
	blogOwner4 := &models.BlogOwner{Introduction: "this is the test4", ImageIconPath: "this is the test"}

	models.DB.Create(blogOwner1)
	models.DB.Create(blogOwner2)
	models.DB.Create(blogOwner3)
	models.DB.Create(blogOwner4)
	result, err := FindPageRecords("1", &models.BlogOwner{})

	if err != nil {
		t.Error(err.Error())
	}

	if s, ok := result.([]models.Category); ok {
		if len(s) != 4 {
			t.Fail()
		}
	}

	models.DB.Delete(blogOwner1)
	models.DB.Delete(blogOwner2)
	models.DB.Delete(blogOwner3)
	models.DB.Delete(blogOwner4)
}
