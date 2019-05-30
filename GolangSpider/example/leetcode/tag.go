package leetcode

import (
	"GolangSpider/GolangSpider/util"
	"github.com/tidwall/gjson"
	"log"
	"strings"
	"time"
)

func SaveTagData()  {
	tagData, err := SaveJsonDataByUrl(tagUrl, tagsJsonfile)
	if err==nil{
		//获取所有的tagSlug信息
		GetAllTagsData(tagData)
	}
}

//爬取所有的Tag数据
func GetAllTagsData(tagData string) {
	cost := util.NewCost(time.Now())
	finish:=0
	slugs,names := ParseTagInfo(tagData)
	tagChan := make(chan DownloadChan)
	for index, slug := range slugs {
		name:=names[index].String()
		if len(name)==0{
			name=slug.String()
		}
		go DownloadTagDataOfType(slug.String(),name, tagChan,)
	}

	for i := 0; i < len(slugs); i++ {
		done := <-tagChan
		if done.Finish {
			log.Println("下载", done.Name, "数据成功")
			finish++
		} else {
			log.Fatalln("下载", done.Name, "数据失败")
		}
	}
	log.Println("所有数据已下载完毕!总共下载：", len(slugs), "\t下载成功：", finish, "\t下载失败：", len(slugs)-finish,
		"\t总耗时：", cost.CostWithNowAsString())
}

//解析tag信息
func ParseTagInfo(data string) ([]gjson.Result,[]gjson.Result) {
	parse := gjson.Parse(data)
	slugs :=parse.Get("topics.#.slug").Array()
	translateName :=parse.Get("topics.#.translatedName").Array()
	return slugs,translateName
}

//下载tag数据
func DownloadTagDataOfType(slug string,name string, done chan DownloadChan) {
	_,err:=SaveTagDataOfType(slug,name)
	if err!=nil{
		done<-DownloadChan{Finish:false,Name:name}
	}else{
		done<-DownloadChan{Finish:true,Name:name}
	}

}

func SaveTagDataOfType(tagSlug string,tagName string) (string,error) {
	data, err := getPostJsonByTemplate(tagDataTemplate, tagData, tagSlug)
	if err==nil{
		util.SaveJsonStr2File(data,strings.Replace(eachTagJsonfile,"#filename#",tagName,-1))
	}
	return data,err
}