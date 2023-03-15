package log

import (
	"log"
	"os"
)

var (
	Info    *log.Logger // 重要的信息
	Warning *log.Logger // 需要注意的信息
	Error   *log.Logger // 非常严重的问题
)

var _asyncWrite bool
var errCh chan []interface{}
var infoCh chan []interface{}
var warnCh chan []interface{}

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
	initInfo()
	initWaring()
	initError()
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

func PrintlnErr(v ...interface{}) {
	if !_asyncWrite {
		Error.Println(v...)
		return
	}
	errCh <- v
}

func PrintlnInfo(v ...interface{}) {
	if !_asyncWrite {
		Info.Println(v...)
		return
	}
	infoCh <- v
}

func PrintlnWarn(v ...interface{}) {
	if !_asyncWrite {
		Warning.Println(v...)
		return
	}
	warnCh <- v
}

func RunAsync() {
	_asyncWrite = true
	infoCh = make(chan []interface{}, 100)
	warnCh = make(chan []interface{}, 100)
	errCh = make(chan []interface{}, 100)
	go func() {
		for {
			select {
			case val := <-infoCh:
				Info.Println(val...)
			case val := <-warnCh:
				Warning.Println(val...)
			case val := <-errCh:
				Error.Println(val...)
			}
		}
	}()
}
