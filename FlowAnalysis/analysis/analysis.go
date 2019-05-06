package analysis

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)


//定义返回客户端json数据的数据结构
type AnalysisParam struct {
	SaveType string `json:"save_type"`
	HandleCount int	`json:"handle_count"`
	ErrorCount	int	`json:"error_count"`
	QPS		int	`json:"qps"`
}

//模拟存储的数据结构
type StoreBlock struct {
	OperType string
	RsType	string//列表页 、还是详情页、 还是首页
	RsId	int//资源Id，即url中的资源Id
	Time 	string
}


//日志数据
type LogInfo struct {
	UId  string
	Time         string
	UserAgent    string
	Url          string
	Refer        string
	IsDetailPage bool
	IsListPage   bool
	IsHome       bool
	RsId	int
}

const (
	DETAIL_PAGE = "/movie/"
	LIST_PAGE   = "/list/"
	SUFFIX=".html"
	STORE_PATH="FlowAnalysis/store.info"
)
//var save *os.File
var exit chan bool
func init() {
	logs.Info("this is the init funciton")
	//save, err := os.OpenFile(STORE_PATH, os.O_CREATE|os.O_APPEND|os.O_RDWR,0766)
	//if err != nil {
	//	panic(err.Error())
	//}
	//isFinish=make(chan bool,1)
	exit =make(chan bool,1 )
}

//1.按行读取日志-->存放在logchannel中
func ReadOneLog(filePath string, logChan chan<- string,isFinish chan bool) {
	logs.Info("ReadOneLog")
	//从文件中按行读取日志
	file, err := os.Open(filePath)
	Check(err)
	defer file.Close()
	rd := bufio.NewReader(file)
	line := ""
	for {
		line, err = rd.ReadString('\n')
		if err==io.EOF{
			//读完了
			isFinish<-true
			time.Sleep(2*time.Second)
			logs.Info("log file is read on the end.")
		}
		Check(err)
		//将读取的日志写入到channel中
		logChan <- line

	}
}

//Check
func Check(e error) {
	if e != nil && e!=io.EOF {
		logs.Error(e.Error())
		panic(e)
	}
}

//2.解析日志信息:数据从logchannel中来，解析完后将解析得到的数据，放进parsedchannel中
func ParseOneLog(logChan <-chan string, parsedChan chan<- LogInfo) {
	logs.Info("ParseOneLog")
	//接收到从logChan中过来的数据，然后处理，处理完塞到parseChan中
	count:=0
	for logStr:=range logChan{
		//从logChan中接收数据
		//data := <-logChan
		//解析日志信息
		logInfo := ParsedLogInfo(logStr)
		//存储到channel中
		parsedChan <- *logInfo
		count++
		//if count%1000==0{
		//	logs.Info(fmt.Sprintf("%#v",logInfo))
		//}

	}
}

//ParsedLogInfo
//解析日志信息
func ParsedLogInfo(data string) *LogInfo {
	logInfo := LogInfo{}
	values, err := url.ParseQuery(data)
	Check(err)
	url := values.Get("url")
	userAgent:=values.Get("ur")
	logInfo.UId=generateUID(url,userAgent)
	logInfo.RsId=getRsId(url)
	logInfo.UserAgent = userAgent
	logInfo.Time = values.Get("time")
	logInfo.Refer = values.Get("refer")
	logInfo.Url = url
	logInfo.IsDetailPage = strings.Contains(url,DETAIL_PAGE )
	logInfo.IsListPage=strings.Contains(url,LIST_PAGE)
	if logInfo.IsDetailPage || logInfo.IsListPage {
		logInfo.IsHome=false
	}else{
		logInfo.IsHome=true
	}
	return &logInfo
}

func getRsId(url string) int {
	rsId:=-1
	var err error
	index1 := strings.Index(url, DETAIL_PAGE)
	if index1!=-1{
		index1=index1+len(DETAIL_PAGE)
		index2:=strings.Index(url,SUFFIX)
		rsId,err=strconv.Atoi(url[index1:index2])
		Check(err)
	}else if index1=strings.Index(url,LIST_PAGE);index1!=1{
		index1=index1+len(LIST_PAGE)
		index2:=strings.Index(url,SUFFIX)
		if index2!=-1{
			rsId,err=strconv.Atoi(url[index1:index2])
			Check(err)
		}
	}
	return rsId
}

//生成uid
func generateUID(url string,userAgent string) string {
	hash := md5.New()
	//io.WriteString(hash,url)
	//io.WriteString(hash,userAgent)
	hash.Write([]byte(url+userAgent))
	uid := hex.EncodeToString(hash.Sum(nil))
	//uid:=hash.Sum(nil)
	return uid
}

//3.统计日志信息：数据从parsedchannel中来，解析完后，将数据存放到statisticschannel中
func StatisticsLog(parsedLog <-chan LogInfo, statChan chan<- StoreBlock) {
	logs.Info("StatisticsLog")
	count :=0
	for parsedInfo:=range parsedLog{
		count++
		//统计数据，然后将其塞入到statChan中，用来存储到数据库中
		//次数可以做uv或者pv的统计，需要注意的是pv需要去重，可以考虑用redis的hyperloglog实现
		storeBlock:=StoreBlock{}
		storeBlock.Time=parsedInfo.Time
		storeBlock.OperType="uv"
		storeBlock.RsId=parsedInfo.RsId
		if parsedInfo.IsDetailPage{
			storeBlock.RsType="movie"
		}else if parsedInfo.IsListPage{
			storeBlock.RsType="list"
		}else{
			storeBlock.RsType="home"
		}
		statChan<-storeBlock
	}
}

//4.存储日志信息：数据从statisticschannel中来，然后将数据存储到redis中
func StoreStatLogInfo(statChan <-chan StoreBlock,isFinish chan bool,save *os.File) {
	blocks:=""
	count:=0
	for{
		select {
			case stblock:=<-statChan:
				blocks=blocks+fmt.Sprintf("%#v",stblock)+"\n"
				count++
				if count%2000==0{
					_, err := save.Write([]byte(blocks))
					fmt.Println(save)
					blocks=""
					if err!=nil{
						logs.Error(err.Error())
						//panic(err.Error())
					}
					//save.Sync()
					logs.Info("writing the ",count,"store block in to file")
					//logs.Info(blocks)
				}
			case <-isFinish:
					save.WriteString(blocks)
					blocks=""
					logs.Info("write the left block into the file")
					exit<-true
		}
	}
	//循环退出，表示完成任务
	//isFinish<-true

}

//5.实时展现存储的数据，借助grafana
func WaitFinish()  {
	if <-exit{
		time.Sleep(3*time.Second)
		logs.Info("the analysis task is finished,you can see the saveinfo file:",STORE_PATH)
	}
}

func GetAnalysisParam() AnalysisParam {
	return AnalysisParam{
		SaveType:    "file",
		HandleCount: 123,
		ErrorCount:0,
		QPS:12242,
	}
}