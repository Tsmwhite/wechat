package validations

import "strings"

// TrimAll 去除所有空格
func TrimAll(content string) string {
	content = strings.ReplaceAll(content, " ", "")
	return content
}

// InputContentFilter 用户输入过滤
func InputContentFilter(content string) string {
	content = strings.ToLower(content)
	content = TrimAll(content)
	content = strings.ReplaceAll(content, "drop", "")
	content = strings.ReplaceAll(content, "select", "")
	content = strings.ReplaceAll(content, "where", "")
	content = strings.ReplaceAll(content, "or", "")
	content = strings.ReplaceAll(content, "=", "")
	content = strings.ReplaceAll(content, "*", "")
	return content
}
