package services

import (
	"log"
)

var s *Service

func init() {
	s = NewService()
}

func isError(err error) bool {
	if err != nil {
		log.Fatal(err.Error())
		return true
	}
	return false
}
