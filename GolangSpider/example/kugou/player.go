package kugou

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
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
	}else if choice=="yes" || choice=="ok" || choice==""{
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
	fmt.Println("暂时未开发该功能，请采用其他播放器播放下载的MV")
	fmt.Println("下载的MV保存路径：",downloadSaveSongDir)

}

func (p *MusicPlayer) PlaySong(cmd *Command) {
	fmt.Printf("%+v\n", cmd)
	//fmt.Println()
	songId,err:=strconv.Atoi(cmd.Arguements[0])
	if err != nil {
		fmt.Println("the songId <",songId,"> you inputed is not corrected.please retry again...")
		return
	}
	if songId>len(downloadSongInfos){
		fmt.Println("the songId<",songId,"> you selected is too large. you should select the song again...")
		return
	}

	if songId<1 {
		fmt.Println("the songId<",songId,"> you selected is not corrected. the songId should >0")
		return
	}

	//以下方式是采用windows自带的播放器播放音乐，行不通
	//后续可采用oto或者其他播放音频的库实现播放音乐，暂时由于系统原因，未实现该部分功能

	//fileName:=songInfos[songId-1].Name
	//fileName="\""+path.Join(downloadSaveSongDir,fileName+".mp3")+"\""
	//fmt.Println(fileName)
	////播放音乐
	//command := exec.Command("cmd", "/c", fileName)
	//command.Stdout=os.Stdout
	//err = command.Run()
	//if err != nil {
	//	fmt.Println("play the song occurs error.",err.Error())
	//}

	fmt.Println("暂时未开发该功能，请采用其他播放器播放下载的音乐")
	fmt.Println("下载的音乐保存路径：",downloadSaveSongDir)
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
			fmt.Fprintf(os.Stderr,"%s\n",err.Error())
		}else{
			songId=songId-1
			if songId>len(songInfos)-1 {
				fmt.Println("当前歌曲编号",songId+1,"过大下载的歌曲总共有",len(songInfos),"首，请等待其他歌曲下载完成后在重新选择歌曲")
				continue
			}
			index++
			song:=songInfos[songId]
			go DownloadMusic(song.FileHash,downloadSaveSongDir,".mp3",songId+1,dmsg)
		}
	}
	download:=0
	hasDownload:=false
	for index>0{
		downloadInfo:=<-dmsg
		hasDownload=true
		if downloadInfo.Success {
			logs.Info("第  (", downloadInfo.FileId, ")  个歌曲  [", downloadInfo.FileName, "]  ", "下载成功")
			download++
		} else {
			logs.Error("第  (", downloadInfo.FileId, ")  个歌曲  [", downloadInfo.FileName, "]  ", "下载失败")
		}
		index--
	}
	if hasDownload{
		logs.Info("指定的歌曲下载完毕，请继续选择操作!!!!")
		logs.Info("总共下载",downloadCount, "个文件!\t下载成功", download,"个文件\t","下载失败",downloadCount-download,"个文件\t","总耗时为", fmt.Sprintf("%v", time.Since(start)))
	}


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
//显示已经下载的歌曲
func (p *MusicPlayer) ShowSong(command *Command) {
	//downloadSaveDir
	ListDownload(downloadSaveSongDir)

}

func  ListDownload(dirPath string) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		logs.Error("read download files error.", err.Error())
		fmt.Println("暂时无法读取到下载的歌曲或者MV信息，请稍后重试...")
		return
	}

	if dirPath==downloadSaveSongDir{
		fmt.Println("歌曲编号\t\t歌曲大小\t\t歌曲名称")
		for index, file := range files {
			fmt.Println("   ", index, "\t\t\t", fmt.Sprintf("%.2f", (float64(file.Size()) / float64(1024*1024))), "M\t\t", file.Name())
			downloadSongInfos=append(downloadSongInfos,&SongInfo{
				FileId:strconv.Itoa(index),
				Name:file.Name(),
				FileSize:fmt.Sprintf("%.2f", (float64(file.Size()) / float64(1024*1024)))+"M",
			})
		}
	}else if dirPath==downloadSaveMVDir{
		fmt.Println("歌曲编号\t\tMV大小\t\tMV名称")
		for index, file := range files {
			fmt.Println("   ", index, "\t\t\t", fmt.Sprintf("%.2f", (float64(file.Size()) / float64(1024*1024))), "M\t\t", file.Name())
			downloadMVInfos=append(downloadMVInfos,&SongInfo{
				FileId:strconv.Itoa(index),
				Name:file.Name(),
				FileSize:fmt.Sprintf("%.2f", (float64(file.Size()) / float64(1024*1024)))+"M",
			})
		}
	}

}
//显示已经下载的MV
func (p *MusicPlayer) ShowMV(command *Command) {
	ListDownload(downloadSaveMVDir)
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
