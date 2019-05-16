package main

import (
	"GolangSpider/util"
	"github.com/astaxie/beego/logs"
)

func main() {

	filename := "baidu.gif"
	url := "http://s1.bdstatic.com/r/www/cache/mid/static/xueshu/img/logo_4b1971d.gif"
	err := util.Download(url, filename)
	if err != nil {
		logs.Info("download file failed:", err.Error())
	} else {
		logs.Info("download file success")
	}
}
