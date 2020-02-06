package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const (
	BETWEEN   = " BETWEEN "
	N_BETWEEN = " NOT BETWEEN "
	IN        = " IN "
	N_IN      = " NOT IN "
	LIKE      = " LIKE "
	EQ        = " = "
	N_EQ      = " != "
	GT        = " > "
	LT        = " < "
	EGT       = " >="
	ELT       = " <= "
	SQL       = "SQL"
)

const (
	AUTO_INCREMENT = "AUTO_INCREMENT"
	PRIMARY        = "PRIMARY"
	UNIQUE         = "UNIQUE"
)

//查询条件
type Condition struct {
	OPERATOR string //运算符 "<",">","=","<=",">=","(NOT)BETWEEN","(NOT)IN","LIKE"
	VALUE    interface{}
	OR       WHERE
}

type WHERE map[string]Condition

func Test() {
	var where, whereOr WHERE
	where = make(map[string]Condition)
	whereOr = make(map[string]Condition)
	whereOr["email"] = Condition{
		OPERATOR: EQ,
		VALUE:    "6766886@qq.com",
	}
	where["member_id"] = Condition{
		OPERATOR: GT,
		VALUE:    100,
		OR:       whereOr,
	}
	where["name"] = Condition{
		OPERATOR: EQ,
		VALUE:    "xiaoming",
	}
	where["mobile"] = Condition{
		OPERATOR: EQ,
		VALUE:    13476223279,
	}

	fmt.Println(where.Analyzing())
}

//查询条件 WHERE 解析
func (w WHERE) Analyzing() string {
	var where string
	for k, v := range w {
		if v.OPERATOR != "" && v.VALUE != nil {
			switch v.OPERATOR {
			case BETWEEN, N_BETWEEN:
				if val, ok := v.VALUE.([]string); ok {
					//string
					where += fmt.Sprintf(" %s %s \"%s\" AND \"%s\" AND", k, v.OPERATOR, val[0], val[1])
				} else if val, ok := v.VALUE.([]int); ok {
					//int
					where += fmt.Sprintf(" %s %s %d AND %d AND", k, v.OPERATOR, val[0], val[1])
				} else if val, ok := v.VALUE.([]int64); ok {
					//int64
					where += fmt.Sprintf(" %s %s %d AND %d AND", k, v.OPERATOR, val[0], val[1])
				} else if val, ok := v.VALUE.([]float64); ok {
					//float64
					where += fmt.Sprintf(" %s %s %f AND %f AND", k, v.OPERATOR, val[0], val[1])
				} else if val, ok := v.VALUE.([]float32); ok {
					//float32
					where += fmt.Sprintf(" %s %s %f AND %f AND", k, v.OPERATOR, val[0], val[1])
				}
			case IN, N_IN:
				if val, ok := v.VALUE.([]string); ok {
					//string
					where += fmt.Sprintf(" %s %s (%s) AND", k, v.OPERATOR, strings.Join(val, ","))
				} else if val, ok := v.VALUE.([]int); ok {
					//int
					str, _ := json.Marshal(val)
					where += fmt.Sprintf(" %s %s (%s) AND", k, v.OPERATOR, string(str))
				} else if val, ok := v.VALUE.([]int64); ok {
					//int64
					str, _ := json.Marshal(val)
					where += fmt.Sprintf(" %s %s (%s) AND", k, v.OPERATOR, string(str))
				} else if val, ok := v.VALUE.([]float64); ok {
					//float64
					str, _ := json.Marshal(val)
					where += fmt.Sprintf(" %s %s (%s) AND", k, v.OPERATOR, string(str))
				} else if val, ok := v.VALUE.([]float32); ok {
					//float32
					str, _ := json.Marshal(val)
					where += fmt.Sprintf(" %s %s (%s) AND", k, v.OPERATOR, string(str))
				}
			case LIKE:
				where += fmt.Sprintf(" %s LIKE %%%s%% AND", k, v.VALUE)
			case GT, LT, EGT, ELT, EQ, N_EQ:
				if _, ok := v.VALUE.(string); ok {
					where += fmt.Sprintf(" %s%s\"%s\" AND", k, v.OPERATOR, v.VALUE)
				} else if _, ok := v.VALUE.(int); ok {
					where += fmt.Sprintf(" %s%s%d AND", k, v.OPERATOR, v.VALUE)
				}
			case SQL:
				where += fmt.Sprintf(" %s AND", v.VALUE)
			}
		}

		if len(v.OR) > 0 {
			fmt.Println("v.OR != nil", v.OR != nil, v.OR, len(v.OR))
			where += " OR (" + v.OR.Analyzing() + ")"
		}
	}

	where = strings.TrimRight(where, "AND")

	return where
}

func Add(table string, data interface{}) (err error) {
	var query string
	var add_data []map[string]reflect.Value
	var fields []string
	var tags map[string]map[string]string
	//反射获取 data type
	typeInfo := reflect.TypeOf(data)
	switch typeInfo.Kind() {
	case reflect.Slice: //slice 时为批量添加
		//获取slice 子类型
		typeInfo = typeInfo.Elem()
		if typeInfo.Kind() != reflect.Struct {
			panic("DATA IS NOT []struct ")
		}
		//获取字段信息
		fields, tags = GetFieldsTags(typeInfo)
		//获取 data 反射值（reflect.Value 类型）
		valInfo := reflect.ValueOf(data)
		//获取slice长度
		len := valInfo.Len()
		for i := 0; i < len; i++ {
			val := valInfo.Index(i)
			//获取 struct 字段值
			res := GetFieldValue(fields, val)
			add_data = append(add_data, res)
		}
	case reflect.Struct:
		//获取字段信息
		fields, tags = GetFieldsTags(typeInfo)
		val := reflect.ValueOf(data)
		res := GetFieldValue(fields, val)
		add_data = append(add_data, res)
	default:
		err = errors.New("func Add() data Type Error")
	}

	//生成sql
	query = insertSql(table, fields, add_data, tags)
	_, err = _DB_.Exec(query)
	//fmt.Println(query,res,err)
	return err
}

func GetInfo(table string, where WHERE) {

}

func GetResult(table string, where WHERE) {

}

func Del(table string, where WHERE) {

}

func Update(table string, data map[string]interface{}, where []WHERE) error {
	var Sql, Update string
	for k, v := range data {
		switch v.(type) {
		case string:
			Update += fmt.Sprintf(" `%s` = \"%s\"", k, v)
		case int, int8, int16, int32, int64:
			Update += fmt.Sprintf(" `%s` = %d ", k, v)
		case float32, float64:
			Update += fmt.Sprintf(" `%s` = %f ", k, v)
		}
		Update += ","
	}
	Update = strings.TrimRight(Update, ",")

	if Update == "" {
		return errors.New("No Update Info")
	}

	Sql = "UPDATE `" + table + "` SET "
	fmt.Println("Update Sql-", Sql)
	return nil
}

//获取结构子字段信息
//param
//@t reflect.Type 反射类型
//return
//@fields []string 表字段slice
//@tags	  map[string]map[string]string 字段标记信息
func GetFieldsTags(t reflect.Type) (fields []string, tags map[string]map[string]string) {
	//初始化tags map
	tags = make(map[string]map[string]string)
	//获取字段数量
	num := t.NumField()
	var field string
	for i := 0; i < num; i++ {
		//struct 字段名
		field = t.Field(i).Name
		fields = append(fields, field)
		tag := make(map[string]string)
		//tag 注明 table 字段名
		tag["field"] = t.Field(i).Tag.Get("json")
		//tag未注明字段时取 struct 字段
		if tag["field"] == "" {
			tag["field"] = field
		}
		//默认值
		tag["default"] = t.Field(i).Tag.Get("default")
		//键
		tag["key"] = t.Field(i).Tag.Get("key")
		tags[field] = tag
	}
	return fields, tags
}

//获取数据反射值 reflect.value
//param
//@fields  []string 表字段slice
//@value   结构体反射值
//return
//将字段名与反射值关联返回map
func GetFieldValue(fields []string, value reflect.Value) map[string]reflect.Value {
	var data map[string]reflect.Value = make(map[string]reflect.Value)
	for i := 0; i < len(fields); i++ {
		data[fields[i]] = value.FieldByName(fields[i])
	}
	return data
}

//生成添加sql
func insertSql(table string, fields []string, values []map[string]reflect.Value, tags map[string]map[string]string) (sql string) {
	var filedStr string
	//valStr 储存值
	valStr := ""
	//遍历添加值
	for i := 0; i < len(values); i++ {
		val := values[i]
		var temp string
		for j := 0; j < len(fields); j++ {
			v := val[fields[j]]
			//tag 标注
			tag := tags[fields[j]]
			//值不存在或（字段为自增并值为0时）
			if !v.IsValid() || (tag["key"] == AUTO_INCREMENT && v.IsZero()) {
				temp += "Null"
			} else {
				//取出对应字段值进行类型判断 拼接字符串
				switch v.Type().Name() {
				case "string":
					temp += "\"" + v.String() + "\""
				case "int", "int32", "int64":
					temp += strconv.FormatInt(v.Int(), 10)
				case "float32", "float64":
					temp += strconv.FormatFloat(v.Float(), 'f', 2, 64)
				default:
					temp += "Null"
				}
			}
			//逗号分割 值
			temp += ","
			if i == (len(values) - 1) {
				filedStr += "`" + tag["field"] + "`,"
			}
		}
		//移除右侧多余逗号
		temp = strings.TrimRight(temp, ",")
		valStr += fmt.Sprintf("(%s)", temp) + ","
	}

	valStr = strings.TrimRight(valStr, ",")
	filedStr = strings.TrimRight(filedStr, ",")
	sql = fmt.Sprintf("INSERT INTO `%s` (%s) VALUES %s", table, filedStr, valStr)
	return
}
