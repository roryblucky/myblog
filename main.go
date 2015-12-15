package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"myblog/models"
	_ "myblog/routers"
	"net/http"
)

func notFound(rw http.ResponseWriter, r *http.Request) {
	errorResult := models.ErrorResult{
		Code:    http.StatusNotFound,
		Message: "Not Found",
	}
	b, _ := json.Marshal(&errorResult)
	buffer := bytes.NewBuffer([]byte{})
	json.HTMLEscape(buffer, b)
	io.WriteString(rw, fmt.Sprintf("%s", buffer))
}

func main() {
	beego.Errorhandler("404", notFound)
	beego.Run()
}
