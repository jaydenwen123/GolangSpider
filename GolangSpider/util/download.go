package util

import (
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
)

//根据url下载图片、视频等二进制文件
func Download(url string, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	//读取响应体的内容
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("download file error!", err.Error())
		return err
	}

	//写入到文件中
	err = ioutil.WriteFile(filename, data, 0755)
	if err != nil {
		logs.Error("download file error!", err.Error())
		return err
	}
	return nil
}
