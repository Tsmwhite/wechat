package res

import (
	"github.com/gin-gonic/gin"
)

const (
	ErrorCodeNormal = iota
	ErrorCodeAuthErr
)

type ErrOption struct {
	Msg       string
	ErrorCode int
}

func ErrCode(c *gin.Context, msg int) {
	Error(c, GetTip(msg))
}

func Success(c *gin.Context, res interface{}) {
	c.JSON(200, gin.H{
		"ok":   true,
		"data": res,
	})
	c.Abort()
}

func Error(c *gin.Context, err error) {
	c.JSON(200, gin.H{
		"ok":       false,
		"msg":      err,
		"err_code": ErrorCodeNormal,
	})
	c.Abort()
}
