package util

import "regexp"

//正则表达式匹配目标
func MatchTarget(regExp, content string) [][]string {
	cp := regexp.MustCompile(regExp)
	//带分组的匹配
	submatchs := cp.FindAllStringSubmatch(content, -1)
	return submatchs
}
