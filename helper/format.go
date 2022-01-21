package helper

import (
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
	return string(runeStr)
}

// Case2Camel 下划线写法转为驼峰写法
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}
