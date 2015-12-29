package utils

import (
	"fmt"
	"github.com/pborman/uuid"
	"io"
	"mime/multipart"
	"myblog/logger"
	"os"
	"path/filepath"
	"strings"
)

func IsEmpty(s string) bool {
	return s == "" || len(s) == 0
}

func GenerateID() string {
	return strings.Replace(uuid.New(), "-", "", -1)
}

func SaveFile(sourceFile multipart.File, fileHeader *multipart.FileHeader, des string) string {
	if sourceFile == nil || fileHeader == nil || IsEmpty(des) {
		return ""
	}
	fileAbsPath := fmt.Sprintf("%s%c%s", des, filepath.Separator, fileHeader.Filename)
	fileLocation := fileAbsPath[:strings.LastIndex(fileAbsPath, string(filepath.Separator))]
	os.MkdirAll(fileLocation, os.ModePerm)
	tempFile, err := os.OpenFile(fileAbsPath, os.O_WRONLY|os.O_CREATE, 0666)
	defer tempFile.Close()
	if err != nil {
		logger.Error("Create temp File failed when do uploading, msg: %s", err.Error())
		return ""
	}
	_, err = io.Copy(tempFile, sourceFile)
	if err != nil {
		logger.Error("Copy File content failed when do uploading, msg: %s", err.Error())
		return ""
	}
	return fileAbsPath
}
