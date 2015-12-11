package services

import (
	"fmt"
	"log"
	"myblog/models"
	"testing"
	"time"
)

var s *Service

func init() {
	s = NewService()
}

func TestAdd(t *testing.T) {
	blog_owner := &models.BlogOwner{ImageIconPath: "This is a test add", Introduction: "This is a test ad"}

	_, err := s.Add(blog_owner)
	if isError(err) {
		t.Fail()
	}

	category := &models.Category{Name: "Java"}
	_, err = s.Add(category)
	if isError(err) {
		t.Fail()
	}

	article := &models.Article{Title: "test add", PostDate: time.Now(), Content: "This is a test add",
		Category: &models.Category{Id: 1}}
	_, err = s.Add(article)
	if isError(err) {
		t.Fail()
	}

	comment := &models.Comment{Content: "This is a test add", Article: article}

	_, err = s.Add(comment)
	if isError(err) {
		t.Fail()
	}
}

func TestRemove(t *testing.T) {
	blog_owner := &models.BlogOwner{ImageIconPath: "This is a test Remove", Introduction: "This is a test Remove"}
	_, err := s.Add(blog_owner)
	if isError(err) {
		t.Fail()
		t.Fatal("insert failed whiled do remove testing")
	}

	_, err = s.Remove(blog_owner)

	if isError(err) {
		t.Fail()
	}
}

func TestUpdate(t *testing.T) {
	blog_owner := &models.BlogOwner{ImageIconPath: "This is a test Update", Introduction: "This is a test Update"}
	_, err := s.Add(blog_owner)

	if isError(err) {
		t.Fail()
		t.Fatal("insert failed whiled do remove testing")
	}

	blog_owner.Introduction = "This is a test Update(updated)"
	_, err = s.Update(blog_owner)

	if isError(err) {
		t.Fail()
	}
}

func TestRead(t *testing.T) {
	article := &models.Article{Id: 1}
	err := s.Read(article)

	if isError(err) {
		t.Fail()
	}

}

func isError(err error) bool {
	if err != nil {
		log.Fatal(err.Error())
		return true
	}
	return false
}
