package validations

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"wechat/helper/format"
)

const TagKey = "validate"
const RuleMultipleSep = " & "
const RuleAssignSep = "="

// 验证类型
const (
	Required   = "required" //默认都是必须值，非必须值用 NoRequired 标记
	NoRequired = "noRequired"
	IsMobile   = "isMobile"
	IsMail     = "isMail"
	Max        = "max"
	Min        = "min"
	MaxLen     = "maxLen"
	MinLen     = "minLen"
	In         = "in"
)

// 方法类型
const (
	Trim   = "trim"
	Filter = "filter"
)

type Rule struct {
	Name, Limit string
}

type RuleModel struct {
	Field      string
	NoRequired bool
	Rules      []Rule
}

type TestStruct struct {
	Mobile    string `validate:"isMobile"`
	Email     string `validate:"isMail & maxLen=50"`
	Number    int    `validate:"max=100 & min=20"`
	Subscribe string `validate:"maxLen=5"`
	Sex       int    `validate:"in=0,1,2"`
	Keyword   string `validate:"filter"`
}

var errorMessages map[string]string

var defaultErrorMessage = map[string]string{
	Required: "missing required parameter %s",
	IsMobile: "%s must be an 11-digit mobile number",
	IsMail:   "%s must be an email address",
	Max:      "the value of %s cannot be greater than %s",
	Min:      "the value of %s is not less than %s",
	MaxLen:   "%s no more than %s characters",
	MinLen:   "%s must be at least %s characters",
	In:       "%s must be one of %s",
}

func Validate(carrier interface{}, message map[string]string, getParam func(string) string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			switch e.(type) {
			case error:
				err = e.(error)
			default:
				//TODO
			}
		}
	}()
	errorMessages = message
	err, ruleModels := ruleParse(carrier)
	if err != nil {
		return err
	}
	params := getParams(carrier, getParam)
	err = verify(ruleModels, params)
	if err != nil {
		return err
	}
	err = assign(carrier, params)
	return err
}

func getParams(carrier interface{}, getParam func(string) string) map[string]string {
	carrierElm := reflect.TypeOf(carrier).Elem()
	fieldsLen := carrierElm.NumField()
	params := make(map[string]string)
	for i := 0; i < fieldsLen; i++ {
		field := carrierElm.Field(i).Name
		params[field] = getParam(format.Camel2Case(field))
	}
	return params
}

func assign(carrier interface{}, params map[string]string) error {
	carrierVal := reflect.ValueOf(carrier).Elem()
	carrierType := reflect.TypeOf(carrier).Elem()
	fieldLen := carrierVal.NumField()
	for i := 0; i < fieldLen; i++ {
		fieldName := carrierType.Field(i).Name
		if params[fieldName] == "" {
			continue
		}
		fieldType := carrierVal.Field(i).Type().Name()
		filed := carrierVal.FieldByName(fieldName)
		switch fieldType {
		case reflect.Int.String():
			value, err := strconv.Atoi(params[fieldName])
			if err != nil {
				return err
			}
			filed.Set(reflect.ValueOf(value))
		case reflect.Int64.String():
			value, err := strconv.ParseInt(params[fieldName], 10, 64)
			if err != nil {
				return err
			}
			filed.Set(reflect.ValueOf(value))
		case reflect.String.String():
			filed.Set(reflect.ValueOf(params[fieldName]))
		}
	}
	return nil
}

func verify(ruleModels []RuleModel, params map[string]string) error {
	for _, model := range ruleModels {
		field := model.Field
		value := params[field]
		// 默认都是必须参数 NoRequired 标识为为非必须
		if value == "" {
			if !model.NoRequired {
				return getErrorMsg(field, Required)
			}
			delete(params, field)
			continue
		}
		for _, rule := range model.Rules {
			switch rule.Name {
			case IsMobile:
				if !format.IsMobile(value) {
					return getErrorMsg(field, IsMobile)
				}
			case IsMail:
				if !format.IsMail(value) {
					return getErrorMsg(field, IsMail)
				}
			case Max, Min:
				if input, err := strconv.Atoi(value); err != nil {
					return err
				} else if val, err := strconv.Atoi(rule.Limit); err != nil {
					return err
				} else {
					if rule.Name == Max && val < input {
						return getErrorMsg(field, Max, rule.Limit)
					} else if rule.Name == Min && val > input {
						return getErrorMsg(field, Min, rule.Limit)
					}
				}
			case MaxLen, MinLen:
				if limit, err := strconv.Atoi(rule.Limit); err != nil {
					return err
				} else {
					valueLen := len(value)
					if rule.Name == MaxLen && limit < valueLen {
						return getErrorMsg(field, MaxLen, rule.Limit)
					} else if rule.Name == MinLen && limit > valueLen {
						return getErrorMsg(field, MinLen, rule.Limit)
					}
				}
			case In:
				arr := strings.Split(rule.Limit, ",")
				ok := func() bool {
					for i := 0; i < len(arr); i++ {
						if arr[i] == value {
							return true
						}
					}
					return false
				}()
				if !ok {
					return getErrorMsg(field, In, rule.Limit)
				}
			case Filter:
				params[field] = InputContentFilter(value)
			case Trim:
				params[field] = TrimAll(value)
			}
		}
	}
	return nil
}

func ruleParse(carrier interface{}) (error, []RuleModel) {
	carrierType := reflect.TypeOf(carrier)
	// carrier 必须是一个struct类型的指针
	if carrierType.Kind() != reflect.Ptr || carrierType.Elem().Kind() != reflect.Struct {
		return errors.New("carrier should be a pointer struct type"), nil
	}
	carrierElm := carrierType.Elem()
	fieldsLen := carrierElm.NumField()
	var resRules []RuleModel
	var field, ruleTag string
	for i := 0; i < fieldsLen; i++ {
		field = carrierElm.Field(i).Name
		ruleTag = carrierElm.Field(i).Tag.Get(TagKey)
		rules := strings.Split(ruleTag, RuleMultipleSep)
		var temp []Rule
		for _, rule := range rules {
			ruleArr := strings.Split(rule, RuleAssignSep)
			r := Rule{
				Name: ruleArr[0],
			}
			if len(ruleArr) > 1 {
				r.Limit = ruleArr[1]
			}
			temp = append(temp, r)
		}
		resRules = append(resRules, RuleModel{
			Field:      field,
			NoRequired: strings.Contains(ruleTag, NoRequired),
			Rules:      temp,
		})
	}
	return nil, resRules
}

func getErrorMsg(field, ruleTag string, limit ...string) error {
	errKey := field
	if ruleTag != Required {
		errKey += strings.Title(ruleTag)
	}
	if message, ok := errorMessages[errKey]; ok == true {
		// 自定义提示
		return errors.New(message)
	} else {
		// 默认提示
		var errStr string
		if len(limit) > 0 {
			errStr = fmt.Sprintf(defaultErrorMessage[ruleTag], field, limit)
		} else {
			errStr = fmt.Sprintf(defaultErrorMessage[ruleTag], field)
		}
		return errors.New(errStr)
	}
}
