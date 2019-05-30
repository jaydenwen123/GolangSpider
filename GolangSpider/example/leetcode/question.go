package leetcode

import (
	"GolangSpider/GolangSpider/util"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var mutex sync.Mutex
var downloadFailed int
//爬取所有的Tag数据
func GetAllQuestions(questionsInfo []*Question) {
	cost := util.NewCost(time.Now())
	finish := 0
	//questionInfo := ParseQuestionInfo(questionsData)
	tagChan := make(chan DownloadChan)
	allFinish := make(chan bool)
	queue:=make([]string,0)
	for index, question := range questionsInfo {
		queue=append(queue,question.QuestionTitleSlug)
		if (index%20==0 && index>0) || index==len(questionsInfo)-1{
			go DownloadQuestion(tagChan,queue)
			queue=make([]string,0)
			time.Sleep(time.Millisecond*30)
		}
		if index == 0 {
			go func() {
				for i := 0; i < len(questionsInfo); i++ {
					done := <-tagChan
					if done.Finish {
						logs.Info("下载", done.Name, "数据成功")
						finish++
					} else {
						logs.Error("下载", done.Name, "数据失败")
					}
				}
				allFinish <- true
			}()
		}
		//time.Sleep(time.Millisecond * 50)
	}

	<-allFinish
	log.Println("所有数据已下载完毕!总共下载：", len(questionsInfo), "\t下载成功：", finish, "\t下载失败：", downloadFailed,"\t重新下载成功：",downloadFailed,
		"\t总耗时：", cost.CostWithNowAsString())
}

//下载所有的题目信息
func DownloadQuestion(done chan DownloadChan,queue []string) {
	for _,slug:=range queue	{
		_, err := SaveQuestionDetail(slug)
		if err == nil {
			done <- DownloadChan{Finish: true, Name: slug}
		} else {
			//done <- DownloadChan{Finish: false, Name: slug}
			log.Println(slug,"下载失败.2秒后将重新下载...")
			mutex.Lock()
			downloadFailed++
			mutex.Unlock()
			time.Sleep(time.Second*2)
			go DownloadQuestion(done,[]string{slug})
		}
		time.Sleep(time.Millisecond*50)
	}
}

//爬取单个题目的详细信息
func SaveQuestionDetail(questionTitleSlug string) (string, error) {
	//替换获取单独题目的详细信息questionParamTemplate中的变量questionParam
	/*param:=strings.Replace(questionParamTemplate,questionParam,questionTitleSlug,-1)
	jsonStr := common.RequestJsonWithPost(commonUrl, headers,param)
	//log.Println(jsonStr)
	if !gjson.Valid(jsonStr){
		log.Fatalln("从服务器拉去题目数据失败，请稍后重试")
		return "",errors.New("从服务器拉去题目数据失败，请稍后重试")
	}*/
	questionDir := strings.Replace(eachQuestionDir, "#filename#", questionTitleSlug, -1)
	util.InitDir(questionDir)
	jsonStr, err := getPostJsonByTemplate(questionParamTemplate, questionParam, questionTitleSlug)
	if err == nil {
		util.SaveJsonStr2File(jsonStr, questionDir+string(os.PathSeparator)+"question.json")
	}
	return jsonStr, err
}


//解析算法题信息
func ParseQuestionInfo(jsonStr string,transJsonStr string) []*Question{
	//从返回的数据中，解析出来
	res := gjson.Parse(jsonStr)
	transRes := gjson.Parse(transJsonStr)
	idRes := res.Get(questionIdPath)
	if idRes.IsArray(){
		questionIdRes:=idRes.Array()
		questions:=make([]*Question,len(questionIdRes))
		questionTitleRes := res.Get(questionTitlePath).Array()
		questionTitleSlugRes := res.Get(questionTitleSlugPath).Array()
		questionTotalAcsRes := res.Get(questionTotalAcsPath).Array()
		questionTotalSubmittedRes := res.Get(questionTotalSubmittedPath).Array()
		questionTotalArticlesRes := res.Get(questionTotalArticlesPath).Array()
		questionDifficultyLevelRes := res.Get(questionDifficultyLevelPath).Array()
		transQuestionTitleRes:=transRes.Get(transQuestionTitlePath).Array()
		tmp:=0
		var ratio string
		for index:=len(questionIdRes)-1;index>=0;index--{
			questionId:=questionIdRes[index]
			calcRatio:=float64(questionTotalAcsRes[index].Int())/float64(questionTotalSubmittedRes[index].Int())*100
			if math.IsNaN(calcRatio){
				ratio="N/A"
			}else{
				ratio=fmt.Sprintf("%.1f",calcRatio)+"%"
			}
			questions[tmp]=&Question{
				QuestionId:questionId.Int(),
				QuestionNo:int64(tmp+1),
				QuestionAC:ratio,
				QuestionTitle:questionTitleRes[index].String(),
				//json数据按照questionId从小到大顺序开始
				TransQuestionTitle:transQuestionTitleRes[tmp].String(),
				QuestionTitleSlug:questionTitleSlugRes[index].String(),
				AnswersCount:questionTotalArticlesRes[index].Int(),
				DifficultLevel:difficlutLevel[int(questionDifficultyLevelRes[index].Int())],
			}
			tmp++
		}
		return questions
	}
	return nil
}


func ShowQuestionList()  {
	//爬取所有算法题列表
	_, questionsInfo := GetAllAlgorithmsList()
	if questionsInfo!=nil{
		fmt.Println(questionsInfo[0].HeadString())
		fmt.Println("---------------------------------------------------------------------------")
		for _,question:=range questionsInfo{
			fmt.Println(question.ToString())
		}
	}
}

func ShowQuestionListWithSize(questionsInfo []*Question,size int) string {

	var res string
	if questionsInfo!=nil{
		res=res+questionsInfo[0].HeadString()+"\n"
		res=res+"-------------------------------------------------------------------------------------------------\n"
		for index,question:=range questionsInfo{
			if index==size{
				break
			}
			res=res+question.ToString()+"\n"
		}
	}
	return res
}


func ShowQuestionInfoWithWeb()  {
	//爬取所有算法题列表
	_, questionsInfo := GetAllAlgorithmsList()
	http.HandleFunc("/questions", func(writer http.ResponseWriter, r *http.Request) {
		res:=ShowQuestionListWithSize(questionsInfo,100)
		io.WriteString(writer,res)
	})

	http.HandleFunc("/question", func(writer http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		ids := r.Form.Get("id")
		fmt.Println(ids)
		if id,err:=strconv.Atoi(ids);err!=nil{
			writer.Write([]byte("参数有误，请重试"))
		}else{
			data, _ := json.MarshalIndent(questionsInfo[id],"","\t")
			detail,_:=SaveQuestionDetail(questionsInfo[id].QuestionTitleSlug)
			var detailFormat bytes.Buffer
			//对未格式化的json字符串格式化处理
			json.Indent(&detailFormat,[]byte(detail),"","\t")
			io.WriteString(writer,"如下是编号为"+fmt.Sprintf("%d",id)+"的题目简要信息：\n")
			writer.Write(data)
			io.WriteString(writer,"\n如下是编号为"+fmt.Sprintf("%d",id)+"题目的详细信息：\n")
			writer.Write(detailFormat.Bytes())
			//fmt.Println(detailFormat.String())
		}
	})

	err := http.ListenAndServe(":1234", nil)
	if err!=nil{
		logs.Error("开启服务器失败....")
	}
}