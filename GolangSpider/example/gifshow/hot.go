package gifshow

import (
	"GolangSpider/GolangSpider/common"
	"GolangSpider/GolangSpider/example/kugou"
	"GolangSpider/GolangSpider/util"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func DownloadWithBatch(epoch int)  {
	cost:=util.NewCost(time.Now())
	batchChan:=make(chan int)
	batchFinish:=make(chan bool)
	if epoch>0{
		for i:=0;i<epoch;i++{
			go DownloadHotVideo(i,batchChan)
			if i==0{
				go func(all int) {
					for tmp:=0;tmp<all;tmp++{
							logs.Info("第",<-batchChan,"批视频下载完成....")
					}
				}(epoch)
			}
			time.Sleep(time.Millisecond*200)
		}
	}
	<-batchFinish
	logs.Info("批量下载视频完成.总共耗时：",cost.CostWithNowAsString())
}

func DownloadHotVideo(id int,batch chan int)  {
	cost:=util.NewCost(time.Now())
	jsonStr := common.RequestJsonWithPost(gifUrl, headers, bodyData)
	//jsonStr:=SendPostRequest(url,headers,bodyData)
	if !gjson.Valid(jsonStr){
		logs.Error("获取json数据失败")
		return
	}
	logs.Info("获取json数据成功")
	filename:="hot_"+fmt.Sprintf("%d",time.Now().Unix())+".json"
	util.SaveJsonStr2File(jsonStr,filename)
	names, urls := ParseVideoInfo(jsonStr)
	if names==nil || urls==nil{
		logs.Error("系统解析json数据出错...")
		return
	}
	downChan:=make(chan kugou.DownloadMsg)
	count:=len(urls)
	finish:=make(chan bool)
	var url string
	for index,urlArr:=range urls{
		name:=names[index].String()
		if urlArr.IsArray() && len(urlArr.Array())==2{
			url=urlArr.Array()[1].String()
		}else if len(urlArr.Array())==1{
			url=urlArr.Array()[0].String()
		}
		go downloadVideo(index,name,url,downChan)
		if index==0{
			logs.Info("正在下载，请等待....")
			go func(downChannel chan kugou.DownloadMsg,count int ) {
				for i:=0;i<count;i++{
					down:=<-downChannel
					if down.Success{
						logs.Info("第",down.FileId,"个视频下载成功：",down.FileName,)
					}else{
						logs.Error("第",down.FileId,"个视频下载失败：",down.FileName,)
					}
				}
				finish<-true
			}(downChan,count)
		}
	}
	<-finish
	logs.Info(count,"个视频下载完毕，总耗时：",cost.CostWithNowAsString())
	batch<-id
}

func downloadVideo(index int,name string, url string,downChan chan kugou.DownloadMsg) {
	err:=util.Download(url,"videos/"+name+".mp4")
	if err==nil{
		downChan<-kugou.DownloadMsg{
			Success:true,
			FileName:name+">>>>"+url,
			FileId:index,
		}
	}else{
		logs.Error("第",index,"个视频下载失败,正在重新下载：",name+">>>>"+url,"")
		time.Sleep(time.Millisecond*200)
		downloadVideo(index,name,url,downChan)
		//downChan<-kugou.DownloadMsg{
		//	Success:false,
		//	FileName:name+">>>>"+url,
		//	FileId:index,
		//}
	}
}

func ParseVideoInfo(jsonStr string) ([]gjson.Result,[]gjson.Result ){
	parse := gjson.Parse(jsonStr)
	shareInfo:=parse.Get("feeds.#.share_info")
	urls:=parse.Get("feeds.#.main_mv_urls.#.url")
	if shareInfo.IsArray() || urls.IsArray(){
		return shareInfo.Array(),urls.Array()
	}
	return nil,nil
}


func SendPostRequest(url string, headers map[string]string, data string) string {
	request,err:=http.NewRequest("POST",url,strings.NewReader(data))
	if err != nil {
		logs.Error(err.Error())
		return ""
	}
	if headers!=nil{
		for key,value:=range headers{
			request.Header.Add(key,value)
		}
	}
	resp, err := http.DefaultClient.Do(request)
	fmt.Println(resp.StatusCode)
	if err!=nil{
		logs.Error("post do error.",err.Error())
		return ""
	}
	defer resp.Body.Close()
	if respData, err := ioutil.ReadAll(resp.Body);err!=nil{
		logs.Error("read data error.",err.Error())
		return ""
	}else{
		return string(respData)
	}

}
