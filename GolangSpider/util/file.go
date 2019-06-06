package util

import "regexp"

var cmp = regexp.MustCompile(`([\/:*?"<>|]+)`)
//去除文件名中的非法字符
func TrimIllegalStr(filename string,replaceStr string) string {
	newStr := cmp.ReplaceAllString(filename, replaceStr)
	return newStr
}