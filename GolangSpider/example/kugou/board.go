package kugou

import (
	"github.com/jaydenwen123/go-util"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/logs"
	"github.com/tidwall/gjson"
	"path"
	"time"
)

func SpiderAllBoardMusic() {
	//1.爬取所有榜单url
	urls := ParseAllBoardUrls()
	//fmt.Println(urls)
	done := make(chan bool, len(urls))
	index := 0
	for key, value := range urls {
		//2.go协程并发的爬取每个榜单的歌曲数据
		//fmt.Println(key,":",value)
		go DownMusicBoardSongs(key, value, done)
		time.Sleep(time.Millisecond*500)
		index++
	}
	//利用channel阻塞程序
	for i := 0; i < len(urls); i++ {
		<-done
		logs.Info("download finish")
		break
	}
}
//下载所有榜单的歌曲
func DownMusicBoardSongs(name string, url string, finish chan bool) {
	//1.创建目录
	savePath:=path.Join(boardSaveDir,name)
	util.InitDir(savePath)
	//fmt.Println("歌曲保存的路径：",savePath)
	//2.解析歌曲数据，歌曲数据放在javascript中，所以需要解析
	songInfos := ParseBoardSongsInfo(url)
	//保存文件
	util.SaveJsonStr2File(songInfos,savePath+"/"+name+".json")
	//3.下载所有的歌曲
	//利用gjson解析歌曲信息，抽取所有的歌曲名称和歌曲hash值
	songHashs:=ExactSongHash(songInfos)
	//4.下载单首歌曲
	done:=make(chan DownloadMsg,len(songHashs))
	for _,hash:=range songHashs{
		go DownloadMusic(hash.String(),savePath,".jpg",0,done)
		time.Sleep(time.Millisecond*500)
	}

	for i:=0;i<len(songHashs);i++{
		<-done
	}

	finish<-true
}

//抽取所有信息中的歌曲hash值
func ExactSongHash(jsonStr string) ([]gjson.Result) {
	hashs := gjson.Get(jsonStr, "#.Hash")
	if hashs.IsArray() {
		return hashs.Array()
	}
	return nil
}

//解析出榜单的歌曲信息，返回的是json字符串
func ParseBoardSongsInfo(url string) string{
	_, data := util.Request(url)
	data = util.MatchStringValue(`global.features =(?s:(.*?))}\]`, data)
	data=data+"}]"
	//得到json数据
	return data
}

//解析所有榜单的url，键存储榜单名称、值存储url
func ParseAllBoardUrls() map[string]string {
	reader := util.ResponseWithReader(boardUrl)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		logs.Error("goquery new document from reader error.",err.Error())
		panic(err.Error())
	}
	//匹配到所有的url之后，遍历
	urls:=make(map[string]string)
	doc.Find(boardPath).Each(func(i int, selection *goquery.Selection) {
		name,_:=selection.Attr("title")
		url,_:=selection.Attr("href")
		urls[name]=url
	})
	return urls
}

//1.爬取所有榜单的url
//2.解析所有榜单中的歌曲信息
//3.保存成文件
//4.根据歌曲的hash值下载歌曲