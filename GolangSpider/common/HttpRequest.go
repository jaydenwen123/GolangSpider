package common

import (
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
)

//通过get发送请求，返回数据
//第一个参数为字节数组，第二个参数为默认编码为utf-8的字符串
func Request(url string) ([]byte, string) {
	//1.发请求，获取数据
	resp, err := http.Get(url)
	if err != nil {
		logs.Error("http get error:", err.Error())
		panic(err.Error())
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("ioutil ReadAll error:", err.Error())
		return nil, ""
	}
	if err = resp.Body.Close(); err != nil {
		logs.Error("resp Body Close error:", err.Error())
		return nil, ""
	}
	return content, string(content)
}
