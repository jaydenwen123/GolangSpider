package leetcode

import "testing"

func TestGetAllAlgorithmsList(t *testing.T) {
	SaveAllAlgorithmsList()
}

func TestGetAllTranslateAlgorithmsList(t *testing.T) {
	_, err := GetAllTranslateAlgorithmsList()
	if err!=nil{
		t.Error(err.Error())
	}
}