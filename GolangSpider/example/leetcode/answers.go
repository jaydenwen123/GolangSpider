package leetcode

import (
	"GolangSpider/GolangSpider/util"
	"os"
)

//爬取单个题目的所有题解信息
func GetAnswersByQuestion(questionSlug string) (string,error){
	questionSaveDir:=questionDir+string(os.PathSeparator)+questionSlug
	util.InitDir(questionSaveDir)
	data, err := getPostJsonByTemplate(answersParamTemplate, answersParam, questionSlug)
	//answersParam
	if err==nil{
		util.SaveJsonStr2File(data,questionSaveDir+string(os.PathSeparator)+"answers.json")
	}
	//answersParamTemplate
	return data,err
}

//爬取单个题目的每条题解详细信息
func GetAnswerDetail(questionSlug,slug string) (string,error) {
	questionSaveDir:=questionDir+string(os.PathSeparator)+questionSlug+string(os.PathSeparator)+answersDir
	util.InitDir(questionSaveDir)
	//answerDetailParam
	//answerDetailParamTemplate
	data, err := getPostJsonByTemplate(answerDetailParamTemplate, answerDetailParam, slug)
	if err == nil {
		util.SaveJsonStr2File(data,questionSaveDir+string(os.PathSeparator)+slug+".json")
	}
	return data,err

}

