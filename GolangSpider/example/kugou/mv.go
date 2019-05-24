package kugou

import (
	"GolangSpider/common"
	"GolangSpider/util"
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/tidwall/gjson"
	"path"
	"strconv"
	"strings"
	"time"
)

//1.搜索MV数据
func SearchMV() []*MVInfo {
	jsonStr := common.RequestJsonWithRetry(mvSearchUrl, HEADER)
	infos:=ParseMVInfo(jsonStr)
	return infos
}

//2.解析MV的关键信息
//采用gjson解析json字符串
func ParseMVInfo(jsonInfo string) []*MVInfo {
	mvNames:=gjson.Get(jsonInfo,"data.lists.#.FileName")
	mvHashs:=gjson.Get(jsonInfo,"data.lists.#.MvHash")
	mvSizes:=gjson.Get(jsonInfo,"data.lists.#.FileSize")
	mvDurations:=gjson.Get(jsonInfo,"data.lists.#.Duration")
	mvHashMarks:=gjson.Get(jsonInfo,"data.lists.#.MvHashMark")
	infos:=make([]*MVInfo,0)
	var mvInfo *MVInfo
	if mvHashs.IsArray() && mvNames.IsArray(){
		hashsArr:=mvHashs.Array()
		namesArr:=mvNames.Array()
		sizesArr:=mvSizes.Array()
		durationsArr:=mvDurations.Array()
		hashMarksArr:=mvHashMarks.Array()
		for index,hash:=range hashsArr{
			key:=generateMVKey(hash.String())
			//1.设置属性值
			mvInfo=&MVInfo{
				MVId:strconv.Itoa(index+1),
				MVName:namesArr[index].String(),
				Hash:hash.String(),
				Key:key,
				Size:sizesArr[index].String(),
				Duration:durationsArr[index].String(),
				HashMark:hashMarksArr[index].String(),
				DetailUrl:generateMVDetailUrl(hash.String(),key),
			}
			infos=append(infos,mvInfo)
		}
	}
	return infos
}

//根据mv的hash值和key值形成url
func generateMVDetailUrl(hash,key string) string {
	url:=strings.Replace(mvInfoUrl,"{hash}",hash,-1)
	url=strings.Replace(url,"{key}",key,-1)
	return url
}

//根据mv的hash值生成key
func generateMVKey(hash string)  string{
	data:=[]byte(strings.ToUpper(hash)+"kugoumvcloud")
	val := md5.Sum(data)
	key:=fmt.Sprintf("%x",val)
	return key
}

//3.下载MV文件
func DownloadSearchMV()  {
	//1.根据接收的关键词，替换url然后请求数据
	searchUrl:=strings.Replace(mvSearchUrl,"{keyword}",keyword,-1)
	searchUrl=strings.Replace(searchUrl,"{pagesize}","1",-1)
	//info:=common.RequestJsonWithRetry(mvSearchUrl,HEADER)
	//2.保存信息
	//获得总共的条数
	searchInfos := common.RequestJson(searchUrl, HEADER)
	total := GetTotal(searchInfos)
	searchUrl=strings.Replace(mvSearchUrl,"{keyword}",keyword,-1)
	searchUrl=strings.Replace(searchUrl,"{pagesize}",total,-1)
	logs.Info("恭喜你，总共搜索到", total, "首MV！！！")
	logs.Info("搜索歌曲的链接：", searchUrl)
	logs.Info("正在搜索数据中，请耐心等待.....")
	searchInfos = common.RequestJsonWithRetry(searchUrl, HEADER)
	//初始化保存歌曲目录
	saveBasePath := path.Join(downloadSaveMVDir, keyword)
	logs.Info("正在初始化目录,请等待......")
	//initSaveDir(saveBasePath)
	initDir(saveBasePath)
	//initSaveDir(downloadSaveSongDir)
	logs.Info("初始化目录完毕.....")
	//解析json数据放在保存前面，采用go协程去解析，解约时间
	//解析json数据，并得到hash
	//3.解析MV信息
	//4.保存MV信息

	//解析json数据放在保存前面，采用go协程去解析，解约时间
	//解析json数据，并得到hash
	if gjson.Valid(searchInfos) {
		logs.Info("正在解析MV数据，请等待........")
		parsed := make(chan bool)
		go func(done chan bool) {
			mvInfos= ParseMVInfo(searchInfos)
			util.Save2JsonFile(mvInfos, saveBasePath+"/mv.json")
			done <- true
		}(parsed)

		//保存json数据到文件中
		filePath := saveBasePath + "/data.json"
		util.SaveJsonStr2File(searchInfos, filePath)
		<-parsed
		logs.Info("解析歌曲数据完毕.......")
		logs.Info("保存歌曲的目录：", saveBasePath)
	}else{
		logs.Error("由于服务器原因，数据获取失败，请重试...")
	}

}

//下载MV文件
func DownloadMV(mv *MVInfo, dirPath string, suffix string, fileIndex int, done chan DownloadMsg) {
	mvInfo := common.RequestJsonWithRetry(mv.DetailUrl, HEADER)
	if gjson.Valid(mvInfo){
		//补全信息
		mv.MVDetail=ParsedMVDetail(mvInfo)
		//下载歌曲信息
		if mvDetail,ok:=mv.MVDetail["sd"];ok{
			//util.Download(mvDetail.MVUrl,dirPath+"/"+mv.MVName+suffix)
			ToDownload(mvDetail, done, mv, fileIndex, dirPath, suffix)
		} else if mvDetail, ok = mv.MVDetail["hd"]; ok {
			ToDownload(mvDetail, done, mv, fileIndex, dirPath, suffix)
		} else if mvDetail, ok = mv.MVDetail["sq"]; ok {
			ToDownload(mvDetail, done, mv, fileIndex, dirPath, suffix)
		} else if mvDetail, ok = mv.MVDetail["rq"]; ok {
			ToDownload(mvDetail, done, mv, fileIndex, dirPath, suffix)
		}
	} else {
		fmt.Println("服务器内部错误，无法下载该MV")
	}

}
//下载一种MV
func ToDownload(mvDetail *MVDetail, done chan DownloadMsg, mv *MVInfo, fileIndex int, dirPath string, suffix string) {
	//打印歌曲信息
	fmt.Println("##########################第"+fmt.Sprintf("%d",fileIndex)+"首MV信息#######################")
	fmt.Println(mv.ToString())
	fmt.Println("MV下载链接：",mvDetail.MVUrl)
	if mvDetail.MVUrl == "" {
		logs.Error("歌曲没有下载链接！！！")
		done <- DownloadMsg{FileName: mv.MVName, FileId: fileIndex, Success: false}
		return
	}
	//3.正式下载
	err := util.Download(mvDetail.MVUrl, dirPath+"/"+mv.MVName+suffix)
	if err != nil {
		//logs.Error("歌曲：",song.Name,"下载失败...")
		done <- DownloadMsg{FileName: mv.MVName, Success: false, FileId: fileIndex}
	} else {
		//logs.Info("歌曲：",song.Name,"下载成功...")
		done <- DownloadMsg{FileName: mv.MVName, Success: true, FileId: fileIndex}
	}
	time.Sleep(time.Millisecond * 50)
}

//解析MVDetail信息
func ParsedMVDetail(mvInfo string) map[string]*MVDetail {
	mvDetails:=make(map[string]*MVDetail,0)
	//mvdata.hd/sd/sq/rq
	//"hash":"aa7aa85082be7ea7625a47e8d390640e",
	//	"filesize":39451791,
	//	"timelength":277757,
	//	"bitrate":1136673,
	//	"downurl":"
	mvDetails["hd"]=generateMVDetail(mvInfo,"hd")
	mvDetails["sd"]=generateMVDetail(mvInfo,"sd")
	mvDetails["sq"]=generateMVDetail(mvInfo,"sq")
	mvDetails["rq"]=generateMVDetail(mvInfo,"rq")
	return mvDetails
}

//MV详细信息
func generateMVDetail(mvInfo string,mvType string) *MVDetail{
	mvdetail:=&MVDetail{}
	mvdetail.MVHash =gjson.Get(mvInfo, "mvdata."+mvType+".hash").String()
	mvdetail.MVSize =gjson.Get(mvInfo, "mvdata."+mvType+".filesize").String()
	mvdetail.Duration =gjson.Get(mvInfo, "mvdata."+mvType+".timelength").String()
	mvdetail.MVUrl =gjson.Get(mvInfo, "mvdata."+mvType+".downurl").String()
	mvdetail.MVHashMark=mvType
	return mvdetail
}
