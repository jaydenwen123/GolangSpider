package gushiwen

import (
	"GolangSpider/GolangSpider/common"
	"GolangSpider/GolangSpider/example/gushiwen/db"
	"GolangSpider/GolangSpider/util"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/logs"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"strings"
)

//1.爬取古诗文网，诗文栏目的所有诗文类型，作者，朝代，形式（诗、词、曲、文言文）
//https://www.gushiwen.org/shiwen/
func SpiderShiwenKindUrl() ( map[string]string, map[string]string, map[string]string, map[string]string ) {
	_, resp := common.Request(SHIWEN_URL)
	//fmt.Println(resp)
	//util.TrimSpace(resp)
	//利用正则表达式解析出所有的分类（类型、作者、朝代、形式）
	targets := util.MatchTarget(shiwenKindRe, resp)
	//fmt.Println(len(targets),len(targets[0]))
	types:=make(map[string]string)
	authors:=make(map[string]string)
	dynastys:=make(map[string]string)
	styles:=make(map[string]string)
	if len(targets)>4 && len(targets)==15 {
		//11:表示类型，
		for index,item :=range targets[11:]{
			//fmt.Println(item)
			//去除空格
			temp:=util.TrimSpace(item[1])
			//fmt.Println(temp)
			dims := util.MatchTarget(singleShiwenRe, temp)
			switch index {
			//类型
			case 0:
				for _,poem:=range dims{
					types[poem[2]]=poem[1]
				}
				//作者
			case 1:
				for _,poem:=range dims{
					authors[poem[2]]=SHIWEN_BASE_URL+poem[1]
				}
				//朝代
			case 2:
				for _,poem:=range dims{
					dynastys[poem[2]]=SHIWEN_BASE_URL+poem[1]
				}
				//形式
			case 3:
				for _,poem:=range dims{
					styles[poem[2]]=SHIWEN_BASE_URL+poem[1]
				}
			}
		}
	}
	return types,authors,dynastys,styles
}
//2.1爬取不同类型诗文下的诗文链接
//<a href="/gushi/guiyuan.aspx">闺怨</a>
func SpiderShiwenByType(url string) []*db.Poem {
	_, resp := common.Request(url)
	targets := util.MatchTarget(typePoemRe, resp)
	curUrl:=""
	poems:=make([]*db.Poem ,0)
	for _,item:=range targets{
		curUrl=item[1]
		if !strings.HasPrefix(curUrl,"http"){
			curUrl=POEM_BASE_URL+curUrl
		}
		poems=append(poems,db.NewPoem(item[2],db.NewAuthor(item[3]),curUrl))
	}
	return poems
}
//2.2爬取不同类型诗文下的诗文链接
//<a href="/gushi/guiyuan.aspx">闺怨</a>
func SpiderShiwenUrlByType(url string) []string {
	_, resp := common.Request(url)
	targets := util.MatchTarget(typePoemRe, resp)
	curUrl:=""
	urls:=make([]string,0)
	for _,item:=range targets{
		curUrl=item[1]
		if !strings.HasPrefix(curUrl,"http"){
			curUrl=POEM_BASE_URL+curUrl
		}
		urls=append(urls,curUrl)
	}
	return urls
}
//3.爬取诗文链接中的内容
//原文、译文、注释、背景、赏析、作者介绍
func SpiderShiwenContent(shiwenUrl string) *db.Poem {
	poem:=&db.Poem{}
	reader := common.ResponseWithReader(shiwenUrl)
	//采用新的方案来解决
	//https://github.com/PuerkitoBio/goquery
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		logs.Error("goquery error.",err.Error())
		panic(err.Error())
	}
	title:=doc.Find("body > div.main3 > div.left > div:nth-child(2) > div.cont > h1").Text()
	//处理掉/
	title=strings.Replace(title,"/","-",-1)
	contentBlock:=doc.Find("div.contson").First()
	content:=contentBlock.Text()
	id,_:=contentBlock.Attr("id")
	id=strings.Replace(id,"contson","",-1)
	//#contson9cee4425b019
	dynasty_author:=doc.Find("body > div.main3 > div.left > div:nth-child(2) > div.cont > p").Text()
	arr:=strings.Split(dynasty_author,"：")
	transloate,like,background:="","",""
	doc.Find("div.contyishang ").Each(func(i int, selection *goquery.Selection) {
			data:=selection.Text()
			//data=strings.TrimSpace(data)
			if strings.Contains(data,"译文及注释"){
				transloate=strings.Replace(util.TrimSpace(data),"译文及注释","",-1)
				transloate=util.TrimSpace(transloate)
			}else if strings.Contains(data,"译文"){
				transloate=strings.Replace(util.TrimSpace(data),"译文","",-1)
				transloate=util.TrimSpace(transloate)
			}else if strings.Contains(data,"赏析"){
				like=strings.Replace(util.TrimSpace(data),"赏析","",-1)
				like=util.TrimSpace(like)
			}else if strings.Contains(data,"创作背景"){
				background=strings.Replace(util.TrimSpace(data),"创作背景","",-1)
				background=util.TrimSpace(background)
			}
	})
	//解析作者信息
	picture,_:=doc.Find("body > div.main3 > div.left > div.sonspic > div.cont > div > a > img").Attr("src")
	profile:=doc.Find("body > div.main3 > div.left > div.sonspic > div.cont > p:nth-child(3)").Text()
	//fmt.Println(title)
	//fmt.Println(content)
	//fmt.Println(dynasty)
	//fmt.Println(transloate)
	//fmt.Println(like)
	//fmt.Println(background)
	poem.ID=id
	poem.Title=title
	poem.Url=shiwenUrl
	poem.Content=content
	if len(arr)>1{
		poem.Dynasty=arr[0]
		poem.Author=db.NewAuthor2(arr[1],profile,picture)
	}else{
		poem.Author=&db.Author{Profile:profile,Picture:picture,Name:arr[0]}
	}
	poem.Translate=transloate
	poem.Like=like
	poem.Background=background
	return poem
}
//4.存储数据到文件或者数据库中
//这里采用PostgreSQL数据库
func DumpPoemToDatabase(poem *db.Poem,dumpType string)  {
	//3.操作
	db.DB.Create(poem)
}
