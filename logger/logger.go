package logger

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var log *logs.BeeLogger

func init() {
	log = logs.NewLogger(1000)
	runmode := beego.AppConfig.String("runmode")
	if runmode == "dev" {
		log.SetLogger("console", "")
	} else {
		log.SetLogger("file", `{"filename":"test.log"}`)
	}
	log.EnableFuncCallDepth(true)
	log.SetLogFuncCallDepth(3)
}

func Debug(format string, params ...interface{}) {
	log.Debug(format, params...)
}

func Info(format string, params ...interface{}) {
	log.Info(format, params...)
}

func Warn(format string, params ...interface{}) {
	log.Warn(format, params...)
}
func Error(format string, params ...interface{}) {
	log.Error(format, params...)
}
