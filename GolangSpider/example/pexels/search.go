package pexels

import (
	"github.com/jaydenwen123/GolangSpider/GolangSpider/example/kugou"
	"github.com/jaydenwen123/go-util"
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/logs"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

//搜索图片,s后为关键字
//https://www.pexels.com/search-photos/?s=muscle
//https://www.pexels.com/search/muscle/?format=js&seed=2019-05-24%2B04%3A35%3A11%2B%2B0000&page=1&type=

func Search()  {
	//计算程序运行时间
	cost := util.NewCost(time.Now())
	//https://www.pexels.com/zh-cn/search/{keyword}/?format=html&page={page}&type=
	//_,data:=util.Request("https://www.pexels.com/search-photos/?s=muscle")
	//1.接收控制台的输入参数：关键字
	lang,searchWord:=AcceptSearchWordFromConsole()
	//2.初始化下载图片的目录
	downloadPath := path.Join(pictureBaseDir, searchWord)
	util.InitDir(downloadPath)
	//2.根据第一页的数据，然后获取总数，计算分页数
	firstPageUrl:=strings.Replace(pageUrl[lang],"{keyword}",searchWord,-1)
	firstPageUrl=strings.Replace(firstPageUrl,"{page}","1",-1)
	logs.Info("第1页数据的url:",firstPageUrl)
	_, resp := util.RequestWithHeader(firstPageUrl,headers)
	//fmt.Println(resp)
	//3.循环请求每页的数据
	var total,pageCount int
	if lang==1{
		total=GetTotalZh(resp)
	}else if lang==2{
		total=GetTotalEn(resp)
	}
		pageCount=GetPageCount(total)
	logs.Info("总共搜索到了",total,"张图片;\t","共",pageCount,"页数据")
	//4.解析每页中图片的参数信息
	pageChan:=make(chan PageContent,pageCount)
	pictureChan:=make(chan Picture,PAGE_SIZE*5)
	pageChan<-PageContent{1,resp}
	for i:=1;i<pageCount;i++{
		go RequestPage(lang,searchWord, i+1,pageChan)
	}

	//5.解析图片信息
	ParsePictures(pageCount,pageChan,pictureChan,downloadPath)

	//6.下载图片
	success, fail := DownloadPicture(total, downloadPath, pictureChan)
	fmt.Println("所有图片下载完毕...","总共下载:",total,"(张)图片\t下载成功:",success,
		"(张)\t下载失败:",fail,"(张)\t总耗时：",cost.CostWithNowAsString())
}

//下载图片
func DownloadPicture(total int ,downloadPath string,pictures chan Picture)(int,int) {
	var savePath string
	var success,fail int
	downloadMsgChan:=make(chan kugou.DownloadMsg,PAGE_SIZE*5)
	for i:=0;i<total;i++{
		picture:=<-pictures
		for key,url:=range picture.DownloadUrls{
			//暂时只下载小的图片
			if  key==SMALL{
				//key==SMALL ||
				savePath= path.Join(downloadPath, picture.Names[key])
				go Download(url,savePath,downloadMsgChan)
			}
		}
	}
	for index:=0;index<total;index++{
		msg:=<-downloadMsgChan
		if !msg.Success{
			logs.Error(msg.FileName,"下载失败")
			fail++
		}else{
			success++
		}
		if index%PAGE_SIZE==0{
			logs.Info("已经下载完成了",index,"张图片,图片保存路径为：",filepath.Dir(msg.FileName))
		}
	}
	return success,fail
}

func Download(url,path string,msgChan chan kugou.DownloadMsg)  {
	err := util.DownloadWithRetry(url, path)
	msg:=kugou.DownloadMsg{}
	msg.FileName=path
	if err!=nil{
		msg.Success=false
	}else{
		msg.Success=true
	}
	msgChan<-msg
	time.Sleep(time.Millisecond*100)
}

//解析图片
func ParsePictures(pageCount int ,pageChan chan PageContent, pictureChan chan Picture,downloadPath string) {
	var pageData PageContent
	for i:=0;i<pageCount;i++{
		pageData=<-pageChan
		go Parse(pageData,pictureChan,downloadPath)
	}
}
//解析每页数据中的图片信息
func Parse(pageContent PageContent,pictureChan chan Picture,downloadPath string) {
	data:=pageContent.Data
	reader, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err!=nil{
		logs.Error("goquery new document by reader error.",err.Error())
	}
	//几个属性
	//data-photo-modal-download-value-small
	//data-photo-modal-download-value-original
	//data-photo-modal-download-value-medium
	//ata-photo-modal-download-value-large
	//data-photo-modal-download-url
	///photo/416809/download/
	//body  div.photos  article
	//获取下载图片的链接(除去?号后面的参数部分)
	//data-photo-modal-image-download-link
	logs.Info("--------------------","第",pageContent.Index,"页数据开始解析","-------------------")
	pictures:=make([]*Picture,0)
	reader.Find("div.photos article").Each(func(i int, selection *goquery.Selection) {
		types := make(map[string]string)
		types[SMALL],_=selection.Attr("data-photo-modal-download-value-small")
		types[ORIGINAL],_=selection.Attr("data-photo-modal-download-value-original")
		types[MEDIUM],_=selection.Attr("data-photo-modal-download-value-medium")
		types[LARGE],_=selection.Attr("data-photo-modal-download-value-large")
		url,_:=selection.Attr("data-photo-modal-download-url")
		id:=util.MatchStringValue(`.*?/photo\D?/(\d+)/`,url)
		link,_:=selection.Attr("data-photo-modal-image-download-link")
		linkSplit:=strings.Split(link,"?")
		//取问号前面的url前缀信息
		baseUrl:=linkSplit[0]
		picture := NewPicture(id,types,baseUrl)
		pictureChan<-*picture
		logs.Info("解析第",pageContent.Index,"页数据","第",i+1,"张图片数据完成")
		pictures=append(pictures,picture)
	})
	logs.Info("--------------------","第",pageContent.Index,"页数据解析完成","-------------------")
	util.Save2JsonFile(pictures,path.Join(downloadPath,"第"+strconv.Itoa(pageContent.Index)+"页.json"))
	logs.Info("--------------------","第",pageContent.Index,"页数据保存完成","-------------------")

}

//请求每页的数据
func RequestPage(lang int,searchWord string, page int,pageChan chan PageContent, ) {
	eachPageUrl := strings.Replace(pageUrl[lang], "{keyword}", searchWord, -1)
	eachPageUrl = strings.Replace(eachPageUrl, "{page}", strconv.Itoa(page), -1)
	logs.Info("第",page,"页数据的url:",eachPageUrl)
	_, resp := util.RequestWithHeader(eachPageUrl,headers)
	pageChan<-PageContent{page,resp}
}

//计算生成总页数
func GetPageCount(total int)  int{
	pageCount:=total/PAGE_SIZE
	if total%PAGE_SIZE!=0{
		pageCount=pageCount+1
	}
	return pageCount
}

//计算总条数
func GetTotalEn(htmlData string) int{
	totalStr := util.MatchStringValue(totalEnRe, htmlData)
	var total float64
	if strings.Contains(totalStr,"K"){
		totalStr=strings.Replace(totalStr,"K","",-1)
		totalStr=strings.TrimSpace(totalStr)
		total,_=strconv.ParseFloat(totalStr,64)
		total=total*1000
		return int(total)
	}else{
		totalStr=strings.TrimSpace(totalStr)
		totalCount,err:=strconv.ParseInt(totalStr,10,64)
		if err!=nil{
			logs.Error("解析总数出错.",err.Error())
		}
		return int(totalCount)
	}
}
//计算总条数
func GetTotalZh(htmlData string) int{
	totalStr := util.MatchStringValue(totalZhRe, htmlData)
	var total float64
	if strings.Contains(totalStr,"千"){
		totalStr=strings.Replace(totalStr,"千","",-1)
		totalStr=strings.TrimSpace(totalStr)
		total,_=strconv.ParseFloat(totalStr,64)
		total=total*1000
		return int(total)
	}else{
		totalStr=strings.TrimSpace(totalStr)
		totalCount,err:=strconv.ParseInt(totalStr,10,64)
		if err!=nil{
			logs.Error("解析总数出错.",err.Error())
		}
		return int(totalCount)
	}
}

//从控制台终端接收搜索的关键词信息
func AcceptSearchWordFromConsole() (int,string) {
	 var searchWord,reply string
	 var language int
	fmt.Printf("%s","请从选择您的语言，暂时支持英文和中文，中文请输入1，英文输入2:")
	 fmt.Scanf("%d\n",&language)
	fmt.Printf("%s","请从控制台输入您的搜索词:")
	//fmt.Scan("%s\n",&searchWord)
	reader := bufio.NewReader(os.Stdin)
	data,_,_:=reader.ReadLine()
	searchWord=string(data)
	lang,ok:=languages[language]
	if !ok{
		lang="未知语言"
	}
	 fmt.Printf("%s%s%s%s%s" ,"您选择的语言是<",lang,">;您的搜索词为:[",searchWord,"]您确定使用该搜索词搜索图片?确定回复yes，否则回复no.")
	 fmt.Scanf("%s\n",&reply)
	 if reply=="" || reply=="yes"{
	 	return language,searchWord
	 }else{
	 	return AcceptSearchWordFromConsole()
	 }
}