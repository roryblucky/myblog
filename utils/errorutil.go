package utils

import "github.com/astaxie/beego"

func LogError(err error) {
	if err != nil {
		beego.Error(err.Error())
	}
}
