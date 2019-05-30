package leetcode

import "fmt"

//题目主要内容
type Question struct {
	//id
	QuestionId 	int64
	//显示的编号
	QuestionNo	int64
	//题名
	QuestionTitle	string
	//翻译后的题名
	TransQuestionTitle string

	//题标题
	QuestionTitleSlug string
	//题解
	AnswersCount	int64
	//通过率
	QuestionAC		string
	//难度
	DifficultLevel	string
}

func (q *Question) ToString() string {
	str:=fmt.Sprintf(" %d",q.QuestionId)+"\t\t\t"+
		fmt.Sprintf(" %d",q.AnswersCount)+"\t\t "+
		q.QuestionAC+"\t\t"+q.DifficultLevel+"\t\t\t"+
		q.TransQuestionTitle+"\t\t"
	return str
}

func (q *Question) HeadString() string {
	return "题目Id\t\t"+"题解\t\t"+"通过率\t\t"+"难度\t\t\t\t"+"题名\t\t"
}


