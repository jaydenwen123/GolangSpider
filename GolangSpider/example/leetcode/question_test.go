package leetcode

import "testing"

func TestGetQuestionDetail(t *testing.T) {

	//two-sum
	_, err := SaveQuestionDetail("two-sum")
	if err!=nil{
		t.Error(err.Error())
	}
	//add-two-numbers
	_,err=SaveQuestionDetail("add-two-numbers")
	if err!=nil{
		t.Error(err.Error())
	}
	//longest-substring-without-repeating-characters
	_,err=SaveQuestionDetail("longest-substring-without-repeating-characters")
	if err!=nil{
		t.Error(err.Error())
	}
}

func TestShowQuestionList(t *testing.T) {
	ShowQuestionList()
}

func TestShowQuestionInfoWithWeb(t *testing.T)  {
	ShowQuestionInfoWithWeb()
}