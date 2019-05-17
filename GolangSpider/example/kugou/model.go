package kugou

type SongInfo struct {
	Name	string
	Img	string
	Url string
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