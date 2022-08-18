package format

import (
	"fmt"
	"hash/crc32"
	"math"
	"reflect"
	"strings"
	"unicode"
)

// Camel2Case 驼峰式写法转为下划线写法
func Camel2Case(name string) string {
	var resRune []rune
	runeStr := []rune(name)
	for i, r := range runeStr {
		if unicode.IsUpper(r) {
			if i != 0 {
				resRune = append(resRune, '_')
			}
			resRune = append(resRune, unicode.ToLower(r))
		} else {
			resRune = append(resRune, r)
		}
	}
	return string(resRune)
}

// Case2Camel 下划线写法转为驼峰写法
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// StructToMap 结构体转map
func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[Camel2Case(t.Field(i).Name)] = v.Field(i).Interface()
	}
	return data
}

// HashCutTable 获取分表名
func HashCutTable(table, uuid string, total int) string {
	hashVal := crc32.ChecksumIEEE([]byte(uuid))
	val := int(hashVal)
	if val > 2147483647 {
		val -= 4294967296
	}
	number := int(math.Abs(float64(val % total)))
	if number < 10 {
		return fmt.Sprintf("%s_0%d", table, number)
	} else {
		return fmt.Sprintf("%s_%d", table, number)
	}
}
