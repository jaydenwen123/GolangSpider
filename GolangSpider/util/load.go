package util

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
)

//从json文件中恢复出golang对象
func LoadObjectFromJsonFile(filePath string,obj interface{}) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		logs.Error("read json file error.", err.Error())
		panic(err.Error())
	}
	err = json.Unmarshal(data, obj)
	if err != nil {
		logs.Error("json to object error.", err.Error())
		panic(err.Error())
	}
}

