package res

import (
	"github.com/gin-gonic/gin"
)

const (
	ErrorCodeNormal = iota + 1
	ErrorCodeAuthErr
)

type ErrOption struct {
	Msg       string
	ErrorCode int
}

func Success(c *gin.Context, res interface{}, message ...string) {
	resMsg := ""
	if len(message) > 0 {
		resMsg = message[0]
	}
	c.JSON(200, gin.H{
		"ok":       true,
		"data":     res,
		"msg":      resMsg,
		"err_code": 0,
		"err_msg":  "",
	})
	c.Abort()
}

func Error(c *gin.Context, err error, code ...int) {
	resCode := ErrorCodeNormal
	if len(code) > 0 {
		resCode = code[0]
	}
	c.JSON(200, gin.H{
		"ok":       false,
		"err_msg":  err.Error(),
		"err_code": resCode,
	})
	c.Abort()
}

func ErrorStr(c *gin.Context, err string, code ...int) {
	resCode := ErrorCodeNormal
	if len(code) > 0 {
		resCode = code[0]
	}
	c.JSON(200, gin.H{
		"ok":       false,
		"err_msg":  err,
		"err_code": resCode,
	})
	c.Abort()
}
