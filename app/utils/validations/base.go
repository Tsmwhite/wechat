package validations

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

const TagKey = "validate"
const MultipleSep = ","
const (
	NoRequired = "noRequired"
	IsMobile   = "isMobile"
	IsMail     = "isMail"
	IsArray    = "isArray"
	Max        = "max"
	Min        = "min"
	In         = "in"
)

type TestStruct struct {
	Mobile    string `validate:"isMobile"`
	Email     string `validate:"isMail,maxLen=50"`
	Number    int    `validate:"max=100,min=20"`
	Subscribe string `validate:"isArray"`
	Sex       int    `validate:"in=1,2"`
}

func Validate(carrier interface{}) error {
	fmt.Println("open")
	carrierType := reflect.TypeOf(carrier)
	if carrierType.Kind() != reflect.Ptr {
		return errors.New("carrier should be a pointer struct type")
	}
	carrierElm := carrierType.Elem()
	if carrierElm.Kind() != reflect.Struct {
		return errors.New("carrier should be a pointer struct type")
	}

	fieldsLen := carrierElm.NumField()
	var field, tag string
	for i := 0; i < fieldsLen; i++ {
		field = carrierElm.Field(i).Name
		fmt.Println("field",field)
		tag = carrierElm.Field(i).Tag.Get(TagKey)
		verify(field, tag)
	}
	return nil
}

func verify(field, ruleTag string) {
	ruleParse(ruleTag)
}

func ruleParse(ruleTag string) {
	rule := strings.Split(ruleTag, MultipleSep)
	for _, v := range rule {
		fmt.Println(v)
	}
}
