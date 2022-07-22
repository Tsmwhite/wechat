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
	workPath, err := os.Getwd()
	if err != nil {
		log.Fatalln("Get work path error:", err)
		return
	}
	logPath := workPath + "/log"
	_, err = os.Stat(logPath)
	if os.IsNotExist(err) {
		err = os.Mkdir(logPath, 0755)
		if err != nil {
			log.Fatalln("Create log dir error:", err)
		}
	}
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
