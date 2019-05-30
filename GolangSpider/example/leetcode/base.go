package leetcode

import (
	"GolangSpider/GolangSpider/common"
	"GolangSpider/GolangSpider/util"
	"github.com/astaxie/beego/logs"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"strings"
)

//封装的统一发送post请求获取json数据
func getPostJsonByTemplate(template string,replace string,keyword string,)  (string, error){
	//1.替换参数
	param:=template
	if len(replace)>0 && len(keyword)>0{
		param=strings.Replace(param,replace,keyword,-1)
	}
	var err error
	jsonStr := common.RequestJsonWithPost(commonUrl, headers, param)
	if !gjson.Valid(jsonStr){
		logs.Error("从服务器拉去题目数据失败，请稍后重试")
		err=errors.New("从服务器拉去数据失败，请稍后重试")
	}
	return jsonStr,err
}

//保存直接根据url获取的json数据
func SaveJsonDataByUrl(url, filename string) (string, error) {
	data := common.RequestJson(url, headers)
	if !gjson.Valid(data) {
		logs.Error("从服务器拉去数据失败，请稍后重试")
		return "", errors.New("从服务器拉去数据失败，请稍后重试")

	}
	util.SaveJsonStr2File(data, filename)
	return data, nil
}

//初始化leetcode目录
func InitLeetcodeDir()  {
	util.InitDir(dataDir)
	util.InitDir(questionDir)
	util.InitDir(recommandDir)
	util.InitDir(tagDir)
}

//用户下载的封装的结构体
type DownloadChan struct {
	Finish bool
	Name 	string
}