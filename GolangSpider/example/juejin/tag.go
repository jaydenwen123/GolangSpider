package juejin

import (
	"github.com/jaydenwen123/go-util"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/tidwall/gjson"
	"strings"
	"time"
)

func GetAllTags(tagUrl string)  {
	url:=strings.Replace(tagUrl,"{page}","1",-1)
	jsonStr := util.RequestJson(url, headers)
	logs.Info("开始关注第",1,"页的标签")
	AddAllTags(jsonStr)
	if gjson.Valid(jsonStr){
		parse:=gjson.Parse(jsonStr)
		total:=parse.Get("d.total").Int()
		pageCount:=int(total/40)
		if total%40!=0{
			pageCount++
		}
		for i:=1;i<pageCount;i++{
			logs.Info("开始关注第",i+1,"页的标签")
			url=strings.Replace(tagUrl,"{page}",fmt.Sprintf("%d",i+1),-1)
			jsonStr = util.RequestJson(url, headers)
			AddAllTags(jsonStr)
		}

	}else{
		logs.Error("occur error.")
	}
}

func AddAllTags(jsonStr string) {
	parse:= gjson.Parse(jsonStr)
	idArr := parse.Get("d.tags.#.id").Array()
	for _,id:=range idArr{
		url:=strings.Replace(addTagUrl,"{id}",id.String(),-1)
		res := util.RequestJsonWithMethod(url, headers, "PUT", "")
		fmt.Println(res)
		time.Sleep(time.Millisecond*50)
	}
}

//保存一个标签的所有页的文章
func SaveAllTagPageArticles()  {
	//1.统计页数
	pageCount := GetPageCount()
	for curPage:=0;curPage<pageCount;curPage++{
		SaveTagPageArticlesAsMarkdown(curPage+1,pageCount)
	}
}
//统计页数
func GetPageCount() int  {
	url:=strings.Replace(tagAllArticlesUrl,"{page}","1",-1)
	url=strings.Replace(url,"{pagesize}","1",-1)
	jsonData := util.RequestJson(url, headers)
	if gjson.Valid(jsonData){
		//解析总条数、计算页数
		parse:=gjson.Parse(jsonData)
		total:=parse.Get("d.total").Int()
		pageCount:=int(total/PAGESIZE)
		if total%PAGESIZE!=0{
			pageCount++
		}
		return pageCount
	}
	return 0
}

//保存一个标签的一页的所有文章
func SaveTagPageArticlesAsMarkdown(page,pageCount int)  {
	url:=strings.Replace(tagAllArticlesUrl,"{page}",fmt.Sprintf("%d",page),-1)
	url=strings.Replace(url,"{pagesize}",fmt.Sprintf("%d",PAGESIZE),-1)
	jsonData := util.RequestJson(url, headers)
	if gjson.Valid(jsonData){
		logs.Info("总共",pageCount,"页，","正在处理第",page,"页的文章")
		articleInfo, _ := ParseArticleInfo(jsonData, tagArticleDetailPath, "")
		if articleInfo!=nil{
			SavePageArticlesAsMarkdown(MARKDOWN_TAG_DIR,page,articleInfo)
		}else{
			logs.Error("解析文章出错")
		}
	}
	
}

