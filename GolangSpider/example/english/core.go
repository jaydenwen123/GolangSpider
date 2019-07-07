package english

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/logs"
	"github.com/jaydenwen123/go-util"
	"io/ioutil"
	"strings"
)

//1.通过主要的url获取到所有的一级栏目链接和二级栏目链接
//2.通过二级栏目链接获取到所有的文章详情链接
//3.保存文章的内容和视频（如果有视频文件的话）

//解析Category信息
func ParseCategory()  ([]*Category,error){
	reader := util.ResponseWithReader(categoryUrl)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		logs.Error("there is occurs error.when build the document by reader.",err.Error())
		return nil,err
	}
	categorys:=make([]*Category,0)
	doc.Find("dl.cl_item").Each(func(index int, selection *goquery.Selection) {
		category:=&Category{}
		//解析标题
		title := selection.Find("dt.cl_title").Text()
		category.Name=title
		category.Channels=make([]*Channel,0)
		//解析二级栏目
		selection.Find("dd a").Each(func(no int, sec *goquery.Selection) {
			name,_:=sec.Attr("title")
			link,_:=sec.Attr("href")
			link=PREFIX+link
			category.Channels=append(category.Channels,&Channel{
				Name:name,
				Link:link,
			})
		})
		//循环记录
		categorys=append(categorys, category)
	})
	return categorys,nil
}

//解析channel的信息
func ParseChannelInfo(category *Category) {
	if category==nil{
		return
	}
	channels:=category.Channels
	for _,channel:=range channels{
		reader := util.ResponseWithReader(channel.Link)
		doc, err := goquery.NewDocumentFromReader(reader)
		if err != nil {
			logs.Error("there is error when create the document.",err.Error())
			continue
		}
		channel.Articles=make([]*Article,0)
		doc.Find("div.contents.frap dl").Each(func(i int, selection *goquery.Selection) {
			//channel下的文章
			id,_:=selection.Attr("id")
			title:=selection.Find("dd.title").Text()
			times:=selection.Find("dd > span.times").Text()
			date:=selection.Find("dd > span.date").Text()
			channel.Articles=append(channel.Articles,&Article{
				Name:title,
				Link:strings.Replace(articleUrlTemplate,"{id}",id,-1),
				Times:times,
				Date:date,
			})
		})

	}
}

func ParseArticleInfo(channel *Channel)  {
	if channel==nil{
		return
	}
	for _,article:=range channel.Articles{
		reader := util.ResponseWithReader(article.Link)
		doc, err := goquery.NewDocumentFromReader(reader)
		if err != nil {
			logs.Error("there is occurs error.",err.Error())
		}
		article.ArticleDetails=make([]*ArticleDetail,0)
		doc.Find("div.contents.frap dl").Each(func(i int, selection *goquery.Selection) {
			title,_:=selection.Attr("title")
			link,_:=selection.Find("a").Attr("href")
			if link=="javascript:void(0)"{
				link=""
			}else{
				link=PREFIX+link
			}
			times:=selection.Find("dd span.times").Text()
			date:=selection.Find("dd span.date").Text()
			article.ArticleDetails=append(article.ArticleDetails,&ArticleDetail{
				Name:title,
				Link:link,
				Date:date,
				Times:times,
			})
		})
	}
}
//解析文章正文的信息
func ParseArticleContent(articleDetail *ArticleDetail)  {
	if articleDetail==nil{
		return
	}
	if articleDetail.Link==""{
		logs.Info("没有文章链接!!!!")
		return
	}
	reader:= util.ResponseWithReader(articleDetail.Link)
	doc, err := goquery.NewDocumentFromReader(reader)
	content:=&Content{}
	if err == nil {
		//先判断是不是视频
		videoLink, ok := doc.Find("div.videoArea video.video").Attr("src")
		if ok{
			logs.Info("is video")
			content.File=videoLink
			content.IsVideo=true
		}else{
			//不是视频则获取音频文件
			htmlText,_:=ioutil.ReadAll(reader)
			audioLink:=util.MatchStringValue(audioLinkRe,string(htmlText))
			content.File=audioLink
			content.IsVideo=false
		}
		fmt.Println(content.File)
		//解析文章段落
		text:=doc.Find("div.article").Text()
		content.Text=text
		content.Name=doc.Find("h1").Text()
	}
	articleDetail.Content=content
}