package kugou

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"os"
	"strconv"
	"time"
)
//播放接口
type Player interface {
	//SEARCH_MV:
	SearchMV(cmd *Command)
	//SEARCH_SONG:
	SearchSong(cmd *Command)
	//LIST_MV:
	ListMV(cmd *Command)
	//LIST_SONG:
	ListSong(cmd *Command)
	//PLAY_MV:
	PlayMV(cmd *Command)
	//PLAY_SONG:
	PlaySong(cmd *Command)
	//DOWNLOAD_SONG:
	DownloadSong(cmd *Command)
	//DOWNLOAD_MV:
	DownloadMV(cmd *Command)
}

type MusicPlayer struct {
}

func (p *MusicPlayer) SearchMV(cmd *Command) {
	fmt.Printf("%+v\n", cmd)
}

func (p *MusicPlayer) SearchSong(cmd *Command) {
	keyword=cmd.Arguement
	fmt.Printf("%+v\n", cmd)
	//1.接收控制台参数
	fmt.Printf("%s <%s> %s","your serach song keyword is:",keyword,"are you sure using this? please select yes or no:")
	var choice string
	fmt.Scanf("%s\n",&choice)
	if choice=="no"{
		keyword=AcceptInputKeyWord()
		//2.下载歌曲
		DownloadSearchMusic()
	}else if choice=="ok" || choice==""{
		//2.下载歌曲
		DownloadSearchMusic()
	}
}

func (p *MusicPlayer) ListMV(cmd *Command) {
	fmt.Printf("%+v\n", cmd)

}

func (p *MusicPlayer) ListSong(cmd *Command) {
	fmt.Printf("%+v\n", cmd)
	//列出所有的歌曲信息
	if songInfos==nil || len(songInfos)==0{
		fmt.Println("暂时没有歌曲，请先搜索歌曲后在执行该操作")
		//time.Sleep(time.Millisecond*100)
		return
	}
	first:=true
	var song *SongInfo
	fmt.Println("\t\t\t",cmd.Arguement,"首歌曲的信息如下：\t\t\t")
	for _,sIndex:=range cmd.Arguements{
		index,_:=strconv.Atoi(sIndex)
		song=songInfos[index-1]
		if first{
			song.PrintTitle()
			first=false
		}
		song.PrintMainInfo()
	}
}

func (p *MusicPlayer) PlayMV(cmd *Command) {

	fmt.Printf("%+v\n", cmd)

}

func (p *MusicPlayer) PlaySong(cmd *Command) {
	fmt.Printf("%+v\n", cmd)

}

func (p *MusicPlayer) DownloadSong(cmd *Command) {
	fmt.Printf("%+v\n", cmd)
	start:=time.Now()
	//遍历下载
	dmsg:=make(chan DownloadMsg)
	index:=0
	downloadCount:=len(cmd.Arguements)
	for _,songIds:=range cmd.Arguements{
		if songId, err := strconv.Atoi(songIds);err!=nil{
			fmt.Fprintf(os.Stderr,"%s",err.Error())
		}else{
			songId=songId-1
			if songId>len(songInfos)-1 {
				fmt.Println("当前歌曲编号",songId+1,"过大下载的歌曲总共有",len(songInfos),"首，请等待其他歌曲下载完成后在重新选择歌曲")
				continue
			}
			index++
			song:=songInfos[songId]
			go DownloadMusic(song.FileHash,downloadSaveDir,".mp3",songId+1,dmsg)
		}
	}
	for index>0{
		downloadInfo:=<-dmsg
		if downloadInfo.Success {
			logs.Info("第  (", downloadInfo.FileId, ")  个歌曲  [", downloadInfo.FileName, "]  ", "下载成功")
		} else {
			logs.Error("第  (", downloadInfo.FileId, ")  个歌曲  [", downloadInfo.FileName, "]  ", "下载失败")
		}
		index--
	}
	logs.Info("指定的歌曲下载完毕，请继续选择操作!!!!")
	logs.Info("总共下载", downloadCount, "个文件!总耗时为", fmt.Sprintf("%v", time.Since(start)))

}
func (p *MusicPlayer) DownloadMV(cmd *Command) {
	fmt.Printf("%+v\n", cmd)
	//遍历下载
	/*dmsg:=make(chan DownloadMsg)
	for _,songIds:=range cmd.Arguements{
		if songId, err := strconv.Atoi(songIds);err!=nil{
			fmt.Fprintf(os.Stderr,"%s",err.Error())
		}else{
			song:=songInfos[songId]
			go DownloadMusic(song.FileHash,likeSaveDir,".mp3",songId,dmsg)
		}
	}
	for  range cmd.Arguements{
		downloadInfo:=<-dmsg
		if downloadInfo.Success {
			logs.Info("第  (", downloadInfo.FileId, ")  个歌曲  [", downloadInfo.FileName, "]  ", "下载成功")
		} else {
			logs.Error("第  (", downloadInfo.FileId, ")  个歌曲  [", downloadInfo.FileName, "]  ", "下载失败")
		}
	}
	logs.Info("指定的歌曲下载完毕，请继续选择操作!!!!")*/
}

//定义命令显示
type Command struct {
	Action     string
	Arguement string
	Arguements []string
}

func NewCommand(action string, arguements []string) *Command {
	return &Command{Action: action, Arguements: arguements}
}
