package router

import "strings"

// "模块名"、"控制器名"、"方法名"规范，全小写，"_"拼接
func PathFormat(s string) string {
	return strings.ReplaceAll(strings.ToLower(strings.TrimSpace(s)), "-", "_")
}
