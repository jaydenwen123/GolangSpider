package leetcode

import (
	"GolangSpider/GolangSpider/util"
)

//初始化目录
func init() {
	InitLeetcodeDir()
}


//func Main() {
//
//	SaveAllAlgorithmsList()
//
//	////爬取单个题目的详细信息
//	//GetQuestionDetail()
//	////爬取单个题目的题解
//	//GetQuestionAnswers()
//}

func SaveAllAlgorithmsList() {
	//爬取所有算法题列表
	jsonStr, questionsInfo := GetAllAlgorithmsList()
	//保存文件
	util.SaveJsonStr2File(jsonStr, algorithmsJsonfile)
	if questionsInfo != nil {
		util.Save2JsonFile(questionsInfo, questionsJsonfile)
	}
	//并发下载所有的题目信息
	GetAllQuestions(questionsInfo)
}




//爬取所有翻译的数据
func GetAllTranslateAlgorithmsList()  (string, error){

	//1.替换参数
	//2.发送请求
	//3.保存数据
	jsonStr, err := getPostJsonByTemplate(transQuestionsParamTemplate, "", "")
	if err==nil{
		util.SaveJsonStr2File(jsonStr,transAlgorithmsJsonfile)
	}
	return jsonStr,err
}

//爬取所有算法题列表
func GetAllAlgorithmsList() (string,[]*Question){
	jsonStr,_:= SaveJsonDataByUrl(algorithmsUrl, algorithmsJsonfile)
	transJsonStr,_:=GetAllTranslateAlgorithmsList()
	questionsInfo := ParseQuestionInfo(jsonStr,transJsonStr)
	return jsonStr,questionsInfo
}

