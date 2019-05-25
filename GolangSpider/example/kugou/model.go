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
	fmt.Println("MV编号\t\t  MV大小\t\t MV时长\t\t\tMV名称\t\t")
	fmt.Println("------------------------------------------------------" +
		"-------------------------------------------------------------------" +
		"-------------------")
}

func (mv *MVInfo) PrintMainInfo() {
	sizeShow := transferFileSize(mv.Size)
	duration:=transferDuration(mv.Duration)
	fmt.Println("  ", mv.MVId, "\t\t ", sizeShow, "\t\t", duration, "\t\t", mv.MVName, "\t\t")
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

func NewSongInfo(FileId string, Name string, AlbumName string, Duration string, FileSize string,FileHash string) *SongInfo {

	if AlbumName=="" || len(AlbumName)==0{
		AlbumName="暂无专辑"
	}
	return &SongInfo{FileId: FileId, Name: Name, AlbumName: AlbumName, Duration: Duration, FileSize: FileSize, FileHash:FileHash}
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
	//fmt.Printf("%3s%2s%18s%15s%15s%50s\n",this.FileId," ",
	//	this.AlbumName,this.FileSize,this.Duration,this.Name)
	sizeShow := transferFileSize(this.FileSize)
	duration:=transferDuration(this.Duration)
	fmt.Println("  ",this.FileId,"\t\t",sizeShow," \t",duration,
		"\t\t\t",this.AlbumName,"\t\t\t\t",this.Name,"\t\t")

}

func (this *SongInfo) PrintTitle()  {
	//fmt.Printf( "%5s%15s%13s%13s%50s\n","歌曲编号","专辑名称","文件大小",
	//	"时长","歌曲名称")
	fmt.Println("歌曲编号\t歌曲大小\t歌曲时长\t\t专辑名称\t\t\t\t\t歌曲名称\t\t")
	fmt.Println("------------------------------------------------------" +
		"-------------------------------------------------------------------" +
		"-------------------")
}



//将整数的文件大小转换为用户能识别的带单位的文件大小
func transferFileSize(size string) string{
	sizeInt, _ := strconv.Atoi(size)
	sizeShow := fmt.Sprintf("%.2f", (float64(sizeInt) / float64(1024*1024))) + "M"
	return sizeShow
}
//将整数转换为分钟和秒的形式
func transferDuration(sDuration string) string {
	duration,_:=strconv.Atoi(sDuration)
	minute:=duration/60
	second:=duration-60*minute
	secondS:=""
	if second<10 && second>0{
		secondS="0"+strconv.Itoa(second)
	}else if second==0{
		secondS="00"
	}else{
		secondS=strconv.Itoa(second)
	}
	return strconv.Itoa(minute)+"m"+secondS+"s"
}
