package validations

import (
	"fmt"
	"testing"
)

func TestValidate(t *testing.T) {
	ts := &TestStruct{}
	Validate(ts)
	fmt.Println("ts",ts)
}