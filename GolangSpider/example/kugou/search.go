package kugou

import (
	"GolangSpider/common"
	"GolangSpider/util"
	"bufio"
	"flag"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/tidwall/gjson"
	"os"
	"path"
	"strings"
	"time"
)

var (
	argMessage=`命令行版的音乐播放器使用说明：
Usage of spider:
  -keyword string
        this is your search keyword!!
`
)

//下载关键词搜索到的歌曲
func DownloadSearchMusic()  {
	//1.接收关键词
	keyword:=AcceptFromConsole()

	//记录下载时间
	start := time.Now()
	//2.发送请求，接收json数据
	search:=strings.Replace(searchUrl,"{}",keyword,-1)
	search=strings.Replace(search,"$","1",-1)
	//获得总共的条数
	searchInfos := common.RequestJson(search,HEADER)
	total:=GetTotal(searchInfos)
	search=strings.Replace(searchUrl,"{}",keyword,-1)
	search=strings.Replace(search,"$",total,-1)
	logs.Info("恭喜你，总共搜索到",total,"首歌曲！！！")
	logs.Info("搜索歌曲的链接：",search)
	searchInfos = common.RequestJson(search,HEADER)
	//初始化保存歌曲目录
	saveBasePath:=path.Join(likeSaveDir,keyword)
	initSaveDir(saveBasePath)
	logs.Info("保存歌曲的目录：",saveBasePath)
	//保存json数据到文件中
	filePath:=saveBasePath+"/data.json"
	util.SaveJsonStr2File(searchInfos,filePath)
	logs.Info("保存搜索到的歌曲数据完毕，路径：",saveBasePath)
	//3.解析json数据，并得到hash
	hashs:=ParseSearchSongsHashs(searchInfos)
	//4.下载歌曲
	DownloadByRank(hashs, saveBasePath)
	//完成时间
	end := time.Now()
	logs.Info("总共下载", downloadMaxCount, "个文件!总耗时为", fmt.Sprintf("%v", end.Sub(start)))
}


//根据歌曲排名，下载前downMaxCount首歌曲
func DownloadByRank(hashs []gjson.Result, saveBasePath string) {
	if downloadMaxCount > len(hashs) {
		downloadMaxCount = len(hashs)
	}
	done := make(chan DownloadMsg, downloadMaxCount)
	finish := make(chan bool)
	for index, hash := range hashs {
		if index >= downloadMaxCount {
			break
		}
		// 并发下载
		go DownloadMusic(hash.String(), saveBasePath, ".mp3", index+1, done)
		if index == 0 {
			logs.Info("---------------正在多线程下载，请耐心等待--------------")
			go func() {
				for i := 0; i < downloadMaxCount; i++ {
					downloadInfo := <-done
					if downloadInfo.Success {
						logs.Info("第  (", downloadInfo.FileId, ")  个文件  [", downloadInfo.FileName, "]  ", "下载成功")
					} else {
						logs.Error("第  (", downloadInfo.FileId, ")  个文件  [", downloadInfo.FileName, "]  ", "下载失败")
					}
				}
				finish <- true
			}()

		}
	}
	<-finish
}

//从所有的歌曲json数据中获取所有的hash值
func ParseSearchSongsHashs(songInfos string) []gjson.Result {
	hashs := gjson.Get(songInfos,"data.lists.#.FileHash")
	if hashs.IsArray(){
		return hashs.Array()
	}
	return nil
}

//从接送文件中获取总的条数
func GetTotal(info string) string {
	return gjson.Get(info,"data.total").String()
}

//从控制台接收参数
func AcceptFromConsole() string {
	var keyword string
	flag.StringVar(&keyword,"keyword","","this is your search keyword!!")
	flag.Parse()
	//用户没有输入参数
	reader:=bufio.NewReader(os.Stdin)
	if keyword=="" || len(keyword)==0{
		fmt.Printf("%s", "please input your keyword:")
		//_, err := fmt.Scanf("%s\n", &keyword)
		line,_,err:=reader.ReadLine()
		keyword=string(line)
		if err != nil {
			logs.Error("accept the keyword error!!!,please try again.",err.Error())
		}
	}
	return keyword
}