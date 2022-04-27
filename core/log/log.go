package log

import (
	"log"
	"os"
)

var (
	Debug   *log.Logger // 记录所有日志
	Info    *log.Logger // 重要的信息
	Warning *log.Logger // 需要注意的信息
	Error   *log.Logger // 非常严重的问题
	DBError *log.Logger // sql问题
)

func init() {
	initDebug()
	initInfo()
	initWaring()
	initError()
}

func initDebug() {
	Debug = initLog("log/debug.log", "TRACE:")
}

func initInfo() {
	Info = initLog("log/info.log", "INFO:")
}

func initWaring() {
	Warning = initLog("log/waring.log", "WARNING:")
}

func initError() {
	Error = initLog("log/error.log", "ERROR:")
}

func initLog(filename, prefix string) *log.Logger {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
		return nil
	}
	return log.New(file, prefix, log.Ldate|log.Ltime|log.Lshortfile)
}
