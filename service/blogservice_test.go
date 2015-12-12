package services

import (
	"log"
	"myblog/models"
	"testing"
)

var s *Service

func init() {
	s = NewService()
}

func TestGetTotalRecords(t *testing.T) {
	blogOwner1 := &models.BlogOwner{Introduction: "this is the test1", ImageIconPath: "this is the test"}
	blogOwner2 := &models.BlogOwner{Introduction: "this is the test2", ImageIconPath: "this is the test"}
	blogOwner3 := &models.BlogOwner{Introduction: "this is the test3", ImageIconPath: "this is the test"}
	blogOwner4 := &models.BlogOwner{Introduction: "this is the test4", ImageIconPath: "this is the test"}

	s.Create(blogOwner1)
	s.Create(blogOwner2)
	s.Create(blogOwner3)
	s.Create(blogOwner4)

	num := s.GetTotalRecords(models.BlogOwner{})
	if num != 4 {
		t.Fail()
		t.Fatal("Actual size is ", num)
	}

	s.Delete(blogOwner1)
	s.Delete(blogOwner2)
	s.Delete(blogOwner3)
	s.Delete(blogOwner4)

}

func TestFindPageRecords(t *testing.T) {
	blogOwner1 := &models.BlogOwner{Introduction: "this is the test1", ImageIconPath: "this is the test"}
	blogOwner2 := &models.BlogOwner{Introduction: "this is the test2", ImageIconPath: "this is the test"}
	blogOwner3 := &models.BlogOwner{Introduction: "this is the test3", ImageIconPath: "this is the test"}
	blogOwner4 := &models.BlogOwner{Introduction: "this is the test4", ImageIconPath: "this is the test"}

	s.Create(blogOwner1)
	s.Create(blogOwner2)
	s.Create(blogOwner3)
	s.Create(blogOwner4)
	page := s.FindPageRecords("3", &models.BlogOwner{})

	for i := page.Records.Front(); i != nil; i = i.Next() {
		if v, ok := i.Value.(models.BlogOwner); ok {
			if models.BlogOwner(v).Introduction != "this is the test3" {
				t.Fail()
			}
		}
	}

	s.Delete(blogOwner1)
	s.Delete(blogOwner2)
	s.Delete(blogOwner3)
	s.Delete(blogOwner4)
}
