package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
	"strings"
	"time"
	"wechat/controller"
	"wechat/model"
	"wechat/socket"
)

type student struct {
	age   int     `json:"age" default:"18"`
	name  string  `json:"name" default:"admin"`
	sex   bool    `json:"sex" default:"1"`
	rank  int64   `json:"rank"`
	socre float64 `json:"socre"`
	money float32
	time  int64
}

func main() {
	salt := controller.GetRandomChar(6)
	var ipStr string
	ip, err_ := controller.GetClientIP()
	if err_ == nil {
		ipStr = ip.String()
	}
	now := time.Now().Unix()
	var data []model.Member
	mem1 := model.Member{
		Name:         "留恋",
		Mobile:       "17777777777",
		Mail:         "17777777777@163.com",
		Salt:         salt,
		Password:     controller.PassWrodMd5("111111", salt),
		PassLook:     "111111",
		Headimg:      "/Uploads/Headimg/Defautl/default_1.jpg",
		RegisterTime: now,
		LoginTime:    now,
		UpdateTime:   now,
		RegisterIp:   ipStr,
		LoginIp:      ipStr,
	}

	mem2 := model.Member{
		Name:         "丽凰",
		Mobile:       "18999999999",
		Mail:         "18999999999@163.com",
		Salt:         salt,
		Password:     controller.PassWrodMd5("111111", salt),
		PassLook:     "111111",
		Headimg:      "/Uploads/Headimg/Defautl/default_1.jpg",
		RegisterTime: now,
		LoginTime:    now,
		UpdateTime:   now,
		RegisterIp:   ipStr,
		LoginIp:      ipStr,
	}
	data = append(data, mem2, mem1)
	err_ = model.Add("z3_member", data)
	fmt.Println("err_", err_)
	return
	fmt.Println("open")
	test1 := student{
		age:  16,
		name: "xiaoming",
		sex:  true,
	}
	//test(test1)
	model.Add("member", test1)
	var test2 []student
	test2 = append(test2, test1)
	test2 = append(test2, student{age: 1, name: "nih", sex: false})
	model.Add("member", test2)
	return
	Start := &socket.Server{
		Addr:   ":8088",
		WSPath: "/ws",
		AuthToken: func(token socket.Token) (id int, err error) {
			fmt.Println(token)
			return
		},
	}
	err := Start.ListenServer()
	fmt.Println(err)

	//注册gin路由 开启端口监听
	ginEngine := gin.Default()
	//router.RegRouter(ginEngine)
	ginEngine.Run(":8088")
}

func test(data interface{}) {
	typeInfo := reflect.TypeOf(data)
	switch typeInfo.Kind() {
	case reflect.Slice:
		typeInfo = typeInfo.Elem()
		fields, _ := GetFieldsTags(typeInfo)
		valInfo := reflect.ValueOf(data)
		len := valInfo.Len()
		var data []map[string]reflect.Value
		for i := 0; i < len; i++ {
			val := valInfo.Index(i)
			res := GetFieldValue(fields, val)
			data = append(data, res)
		}
		fmt.Println("data", data)
		insertSql("member", fields, data)
	case reflect.Struct:
		//获取字段信息
		//fields,tags := GetFieldsTags(typeInfo)
		valInfo := reflect.ValueOf(data)
		fmt.Println("val222", valInfo)
	default:
		panic("Type Error")
	}
}

func GetFieldsTags(t reflect.Type) (fields []string, tags map[string]map[string]string) {
	tags = make(map[string]map[string]string)
	//获取字段数量
	num := t.NumField()
	var field string
	for i := 0; i < num; i++ {
		//tag 注明字段名
		field = t.Field(i).Tag.Get("json")
		if field == "" {
			field = t.Field(i).Name
		}
		fields = append(fields, field)
		tag := make(map[string]string)
		tag["default"] = t.Field(i).Tag.Get("default")
		tags[field] = tag
	}
	return fields, tags
}

func GetFieldValue(fields []string, value reflect.Value) map[string]reflect.Value {
	var data map[string]reflect.Value = make(map[string]reflect.Value)
	for i := 0; i < len(fields); i++ {
		val := value.FieldByName(fields[i])
		data[fields[i]] = val
	}
	return data
}

//生成添加sql
func insertSql(table string, fields []string, values []map[string]reflect.Value) (sql string) {
	fstr := strings.Join(fields, "`,`")
	valStr := ""
	for i := 0; i < len(values); i++ {
		val := values[i]
		var temp string
		for j := 0; j < len(fields); j++ {
			v := val[fields[j]]
			fmt.Println("type mes", v.Type().String())
			switch v.Type().String() {
			case "string":
				temp += "\"" + v.String() + "\""
			case "int", "int32", "int64":
				temp += strconv.FormatInt(v.Int(), 10)
			case "float32", "float64":
				temp += strconv.FormatFloat(v.Float(), 'f', 2, 64)
			default:
				temp += "Null"
			}
			temp += ","
		}

		temp = strings.TrimRight(temp, ",")
		valStr += fmt.Sprintf("(%s)", temp) + ","
	}
	valStr = strings.TrimRight(valStr, ",")
	sql = fmt.Sprintf("INSERT INTO `%s` VALUES (`%s`) %s", table, fstr, valStr)
	return
}
