package utils

import (
	"github.com/pborman/uuid"
	"strings"
)

func IsEmpty(s string) bool {
	return s == "" || len(s) == 0
}

func GenerateID() string {
	return strings.Replace(uuid.New(), "-", "", -1)
}
