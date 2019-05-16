package util

import (
	"github.com/astaxie/beego/logs"
	"regexp"
	"strconv"
)

//正则表达式去除空格
func TrimSpace(item string) string {
	cp, _ := regexp.Compile(`\s+`)
	s := cp.ReplaceAllString(item, "")
	return s
}

//正则表达式匹配目标
func MatchTarget(regExp, content string) [][]string {
	cp := regexp.MustCompile(regExp)
	//带分组的匹配
	submatchs := cp.FindAllStringSubmatch(content, -1)
	return submatchs
}

//正则表达式匹配单个字符串
//常用的场景有如下几种：
//<h1 onclick="GotoUrl\('/realtime'\)" class="trtop">safsadfsadfsad</h1>
//<h1 onclick="GotoUrl\('/realtime'\)" class="trtop">(.*?)</h1>

//<h6><span id="week"></span>safasdfsdaf</h6>
//<h6><span id="week"></span>(.*?)</h6>

//<h1 onclick="GotoUrl\('/comming'\)" class="trtop">asfasfsad</h1>
//<h1 onclick="GotoUrl\('/comming'\)" class="trtop">(.*?)</h1>

func MatchStringValue(regExp, content string) string {
	cp := regexp.MustCompile(regExp)
	//带分组的匹配
	submatchs := cp.FindAllStringSubmatch(content, -1)
	return submatchs[0][1]
}

//正则表达式匹配单个int类型的值
//常用的场景有如下几种：
//<h1 onclick="GotoUrl\('/realtime'\)" class="trtop">123</h1>
//<h1 onclick="GotoUrl\('/realtime'\)" class="trtop">(.*?)</h1>

//<h6><span id="week"></span>2342</h6>
//<h6><span id="week"></span>(.*?)</h6>

//<h1 onclick="GotoUrl\('/comming'\)" class="trtop">445</h1>
//<h1 onclick="GotoUrl\('/comming'\)" class="trtop">(.*?)</h1>
func MatchIntValue(regExp, content string) int64 {
	target,err:=strconv.ParseInt(MatchStringValue(regExp,content),10,64)
	if err != nil {
		logs.Error(" string parse bool failed:",err.Error())
		return 0
	}
	return target
}

//正则表达式匹配单个float类型的值
//常用的场景有如下几种：
//<h1 onclick="GotoUrl\('/realtime'\)" class="trtop">123.234</h1>
//<h1 onclick="GotoUrl\('/realtime'\)" class="trtop">(.*?)</h1>

//<h6><span id="week"></span>234.2</h6>
//<h6><span id="week"></span>(.*?)</h6>

//<h1 onclick="GotoUrl\('/comming'\)" class="trtop">4.45</h1>
//<h1 onclick="GotoUrl\('/comming'\)" class="trtop">(.*?)</h1>
func MatchFloatValue(regExp, content string) float64 {
	target, err := strconv.ParseFloat(MatchStringValue(regExp, content),64)
	if err != nil {
		logs.Error(" string parse bool failed:", err.Error())
		return 0.0
	}
	return target
}

//正则表达式匹配单个bool类型的值
//常用的场景有如下几种：
//<h1 onclick="GotoUrl\('/realtime'\)" class="trtop">false</h1>
//<h1 onclick="GotoUrl\('/realtime'\)" class="trtop">(.*?)</h1>

//<h6><span id="week"></span>true</h6>
//<h6><span id="week"></span>(.*?)</h6>

//<h1 onclick="GotoUrl\('/comming'\)" class="trtop">false</h1>
//<h1 onclick="GotoUrl\('/comming'\)" class="trtop">(.*?)</h1>
func MatchBoolValue(regExp, content string) bool {
	target,err:=strconv.ParseBool(MatchStringValue(regExp,content))
	if err != nil {
		logs.Error(" string parse bool failed:",err.Error())
		return false
	}
	return target
}

