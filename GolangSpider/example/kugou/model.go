package kugou

import (
	"fmt"
	"strconv"
)

type MVDetail struct {
	MVHash   string
	MVKey      string
	MVHashMark string
	Duration   string
	MVSize     string
	MVUrl	string
}

//mv信息
type MVInfo struct {
	MVId	string
	MVName 	string
	MVDetail map[string]*MVDetail
	Hash 	string
	Key		string
	DetailUrl	string
	HashMark string
	Duration   string
	Size     string
}

func (mv *MVInfo) PrintTitle() {
	fmt.Println("MV编号\t\t  MV大小\t\tMV时长\t\t\tMV名称\t\t")
	fmt.Println("------------------------------------------------------" +
		"-------------------------------------------------------------------" +
		"-------------------")
}

func (mv *MVInfo) PrintMainInfo() {
	fmt.Println("  ",mv.MVId,"\t\t",mv.Size,"\t\t",mv.Duration,"\t\t",mv.MVName,"\t\t")
}

func (mv *MVInfo) ToString() string{
	return "MV名称："+mv.MVName+"\nMV原链接："+mv.DetailUrl+"\n该MV总共有:"+strconv.Itoa(len(mv.MVDetail))+"个版本"
}


type SongInfo struct {
	FileId string
	//歌曲名称
	Name	string
	//专辑名称
	AlbumName	string
	//时长
	Duration  string
	//文件大小
	FileSize string

	//歌曲Hash
	FileHash  string

	//图片链接
	Img	string
	//歌曲下载链接
	Url string
	//获取歌曲json数据的链接地址
	SourceUrl string
}






//下载文件信息
type DownloadMsg struct {
	FileId	int
	FileName	string
	Success		bool
	DownloadUrl	string
}
//打印歌曲的信息
func (this *SongInfo) ToString() string {
	return "歌曲名称："+this.Name+ "\n歌曲下载路径："+this.Url+"\n歌曲原路径："+this.SourceUrl
}

//歌曲信息
func (this *SongInfo) PrintMainInfo() {
	fmt.Printf("%3s%2s%18s%15s%15s%50s\n",this.FileId," ",
		this.AlbumName,this.FileSize,this.Duration,this.Name)
}

func (this *SongInfo) PrintTitle()  {
	fmt.Printf( "%5s%15s%13s%13s%50s\n","歌曲编号","专辑名称","文件大小",
		"时长","歌曲名称")
}


