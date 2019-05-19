package kugou

import (
	"GolangSpider/common"
	"GolangSpider/util"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/tidwall/gjson"
	"os"
	"strings"
	"time"
)

//目录不存在则创建，存在则跳过
func initSaveDir(basePath string) {
	_, err := os.Stat(basePath)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(basePath,os.ModeDir)
		if err != nil {
			logs.Error("create save board music direcotry error.")
		}
	}

}

//根据每首歌曲的hash值下载歌曲
func DownloadMusic(hash,savePath string,fileSuffix string,fileIndex int,done chan DownloadMsg) {
	//1.首先根据hash值获取歌曲的json数据
	songUrl:=strings.Replace(songInfoTemplateUrl,"{}",hash,-1)
	songJson := common.RequestJson(songUrl, HEADER)
	//估计是cookie变化了，所以需要重新设置一次cookie
	//time.Sleep(time.Millisecond*500)
	//2.提取歌曲下载链接
	//audio_name,img,play_url
	song := ExactSongInfo(songJson)
	song.SourceUrl=songUrl
	//打印歌曲信息
	fmt.Println("##########################第"+fmt.Sprintf("%d",fileIndex)+"首歌曲信息#######################")
	fmt.Println(song.ToString())
	//fmt.Println()
	if song.Url==""{
		logs.Error("歌曲没有下载链接！！！")
		done<-DownloadMsg{FileName:song.Name,FileId:fileIndex,Success:false}
		return
		//time.Sleep(time.Millisecond*20)
		//DownloadMusic(hash,savePath,done)
		//return
	}
	//3.正式下载
	err := util.Download(song.Url, savePath+"/"+song.Name+fileSuffix)
	if err!=nil{
		//logs.Error("歌曲：",song.Name,"下载失败...")
		done<-DownloadMsg{FileName:song.Name,Success:false,FileId:fileIndex}
	}else{
		//logs.Info("歌曲：",song.Name,"下载成功...")
		done<-DownloadMsg{FileName:song.Name,Success:true,FileId:fileIndex}
	}
	time.Sleep(time.Millisecond*50)
}

//提取歌曲的重要信息
func ExactSongInfo(songJson string) *SongInfo {
	song:=&SongInfo{}
	song.Name = gjson.Get(songJson, SONG_NAME_PATH).String()
	song.Img = gjson.Get(songJson, IMG_URL_PATH).String()
	song.Url = gjson.Get(songJson, PLAY_URL_PATH).String()
	return song
}

