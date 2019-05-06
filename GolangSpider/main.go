package main

import (
	"GolangSpider/util"
	"github.com/astaxie/beego/logs"
)

func main() {

	filename := "baidu.html"
	url := "http://www.baidu.com"
	err := util.Download(url, filename)
	if err != nil {
		logs.Info("download file failed:", err.Error())
	} else {
		logs.Info("download file success")
	}
	filename = "taobao.html"
	url = "http://taobao.com"
	err = util.Download(url, filename)
	if err != nil {
		logs.Info("download file failed:", err.Error())
	} else {
		logs.Info("download file success")
	}
}
