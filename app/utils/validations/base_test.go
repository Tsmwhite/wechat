package validations

import (
	"fmt"
	"testing"
)

func TestValidate(t *testing.T) {
	ts := &TestStruct{}
	err := Validate(ts, map[string]string{
		"MobileIsMobile": "手机号码格式不正确",
	}, func(s string) string {
		data := map[string]string{
			"mobile":    "11888888888",
			"email":     "15888888888@163.com",
			"number":    "99",
			"subscribe": "1,2,3",
			"sex":       "0",
		}
		if res, ok := data[s]; ok {
			return res
		}
		return ""
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println("ts", ts)
}
